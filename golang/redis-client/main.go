package main

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var redisAddrs = os.Getenv("REDIS_ADDR")

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("module", "redis-client")

	log.Info("start")

	ctx := context.TODO()

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: strings.Split(redisAddrs, ","),
	})

	key := "key"

	value := time.Now().Format(time.RFC3339)
	if err := client.SetEx(ctx, key, value, time.Second*3).Err(); err != nil {
		log.Error("SetEx failed", slog.Any("error", err))
		os.Exit(1)
	}
	log.Info("SetEx success")

	time.Sleep(time.Second * 2)

	if value, err := client.Get(ctx, key).Result(); err != nil {
		if err == redis.Nil {
			log.Info("Get (after 2sec) Not found")
		} else {
			log.Error("Get (after 2sec) failed", slog.Any("error", err))
			os.Exit(1)
		}
	} else {
		log.Info("Get (after 2sec) success", slog.Any("value", value))
	}

	time.Sleep(time.Second * 2)

	if value, err := client.Get(ctx, key).Result(); err != nil {
		if err == redis.Nil {
			log.Info("Get (after 4sec) Not found")
		} else {
			log.Error("Get (after 4sec) failed", slog.Any("error", err))
			os.Exit(1)
		}
	} else {
		log.Info("Get (after 4sec) success", slog.Any("value", value))
	}

	os.Exit(0)
}
