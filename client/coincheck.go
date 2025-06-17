package client

import (
	"fmt"
	"os"

	"github.com/sg3t41/go-coincheck/pkg/coincheck"
	"github.com/sg3t41/coincheck-mcp/config"
)

type Client struct {
	*coincheck.Coincheck
}

func NewWithConfig(cfg *config.Config) (*Client, error) {
	client, err := coincheck.New(
		coincheck.WithCredentials(cfg.APIKey, cfg.APISecret),
		coincheck.WithREST(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create coincheck client: %w", err)
	}

	return &Client{client}, nil
}

func NewWithEnv() (*Client, error) {
	apiKey := os.Getenv("COINCHECK_API_KEY")
	apiSecret := os.Getenv("COINCHECK_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("COINCHECK_API_KEY and COINCHECK_API_SECRET environment variables are required")
	}

	client, err := coincheck.New(
		coincheck.WithCredentials(apiKey, apiSecret),
		coincheck.WithREST(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create coincheck client: %w", err)
	}

	return &Client{client}, nil
}