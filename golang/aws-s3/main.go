package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("module", "aws-s3")

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	client := s3.NewFromConfig(cfg)

	if output, err := client.ListBuckets(ctx, &s3.ListBucketsInput{}); err != nil {
		log.Error(err.Error())
	} else {
		for _, bucket := range output.Buckets {
			log.Info(*bucket.Name)
		}
	}
}
