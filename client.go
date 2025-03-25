package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hse-telescope/logger"
)

// NewClient ...
func NewClient(ctx context.Context, config Config) (*s3.Client, error) {
	awsCfg, err := config.LoadAwsConfig(ctx)
	if err != nil {
		logger.Error(ctx, "failed to load aws config", err)
		return nil, err
	}
	return s3.NewFromConfig(awsCfg), nil
}

// NewPresignedClient ...
func NewPresignedClient(ctx context.Context, config Config) (*s3.PresignClient, *s3.Client, error) {
	client, err := NewClient(ctx, config)
	if err != nil {
		return nil, nil, err
	}
	return s3.NewPresignClient(client), client, nil
}
