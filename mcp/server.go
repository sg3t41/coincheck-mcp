package mcp

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sg3t41/go-coincheck/pkg/coincheck"
	"github.com/sg3t41/coincheck-mcp/config"
)

// Server represents the MCP server with Coincheck client
type Server struct {
	client *coincheck.Coincheck
}

// NewServerWithConfig creates a new MCP server with provided config
func NewServerWithConfig(cfg *config.Config) (*Server, error) {
	client, err := coincheck.New(
		coincheck.WithCredentials(cfg.APIKey, cfg.APISecret),
		coincheck.WithREST(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create coincheck client: %w", err)
	}

	return &Server{client: client}, nil
}

// NewServer creates a new MCP server using environment variables
func NewServer() (*Server, error) {
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

	return &Server{client: client}, nil
}

// Run starts the MCP server and handles incoming requests
func (s *Server) Run() {
	// 読み取り用
	decoder := json.NewDecoder(os.Stdin)
	// 書き込み用
	encoder := json.NewEncoder(os.Stdout)

	for {
		var req Request
		if err := decoder.Decode(&req); err != nil {
			log.Printf("Error decoding request: %v", err)
			continue
		}

		resp := s.HandleRequest(req)
		// Only send response for requests, not notifications
		if resp.Jsonrpc != "" {
			if err := encoder.Encode(resp); err != nil {
				log.Printf("Error encoding response: %v", err)
			}
		}
	}
}