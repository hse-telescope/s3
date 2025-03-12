package s3

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewClient ...
func NewClient(ctx context.Context, config Config) (*s3.Client, error) {
	awsCfg, err := config.LoadAwsConfig(ctx)
	if err != nil {
		logger.Error("failed to load aws config", err)
		return nil, err
	}
	client := s3.NewFromConfig(awsCfg)
	if client == nil {
		logger.Error("failed to load aws client")
		return nil, errors.New("TODO")
	}
	return client, nil
}

// NewPresignedClient ...
func NewPresignedClient(ctx context.Context, config Config) (*s3.PresignClient, *s3.Client, error) {
	client, err := NewClient(ctx, config)
	if err != nil {
		return nil, nil, err
	}
	presigned := s3.NewPresignClient(client)
	if presigned == nil {
		logger.Error("failed to load aws presigned client")
		return nil, nil, errors.New("TODO")
	}
	return presigned, client, nil
}
