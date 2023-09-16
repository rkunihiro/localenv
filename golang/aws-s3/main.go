package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var endpoint = os.Getenv("AWS_ENDPOINT")

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("module", "aws-s3")

	ctx := context.TODO()
	optFuns := []func(*config.LoadOptions) error{
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if endpoint != "" {
					return aws.Endpoint{
						URL:           endpoint,
						SigningRegion: "ap-northeast-1",
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		)),
	}
	cfg, err := config.LoadDefaultConfig(ctx, optFuns...)
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
