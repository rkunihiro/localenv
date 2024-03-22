package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       uint64    `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Created  time.Time `db:"created" json:"created"`
	Modified time.Time `db:"modified" json:"modified"`
}

var dsn = os.Getenv("DATABASE_DSN")

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("module", "db-client")

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Error("sqlx.Open failed", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()

	st, err := db.PrepareNamedContext(context.Background(), "SELECT id,name,created,modified FROM User ORDER BY id ASC")
	if err != nil {
		log.Error("db.PrepareNamedContext failed", slog.Any("error", err))
		os.Exit(1)
	}
	defer st.Close()

	rows, err := st.QueryxContext(context.Background(), map[string]any{})
	if err != nil {
		log.Error("st.QueryxContext failed", slog.Any("error", err))
		os.Exit(1)
	}
	for rows.Next() {
		user := &User{}
		if err := rows.StructScan(user); err != nil {
			log.Error("rows.StructScan failed", slog.Any("error", err))
			os.Exit(1)
		}
		log.Info("rows.StructScan success", slog.Any("user", user))
	}
}
