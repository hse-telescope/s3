package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// ItemConfig ...
type ItemConfig struct {
	Bucket    string        `yaml:"bucket"`
	KeySuffix string        `yaml:"key_suffix"`
	Expires   time.Duration `yaml:"expires"`
}

// ItemsConfig ...
type ItemsConfig struct {
	Items map[string]ItemConfig `yaml:"items"`
}

// Config ...
type Config struct {
	URL             string `yaml:"url"`
	SecretAccessKey string `yaml:"secret_access_key"`
	AccessKeyID     string `yaml:"access_key_id"`
}

// LoadAwsConfig ...
func (c *Config) LoadAwsConfig(ctx context.Context) (aws.Config, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: c.URL,
		}, nil
	})
	return config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     c.AccessKeyID,
				SecretAccessKey: c.SecretAccessKey,
			},
		}),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithRegion("auto"),
	)
}
