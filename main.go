package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sg3t41/go-coincheck/pkg/coincheck"
)

type MCPServer struct {
	client *coincheck.Coincheck
}

type MCPRequest struct {
	Jsonrpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      interface{}            `json:"id"`
}

type MCPResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"inputSchema"`
}

func NewMCPServer() (*MCPServer, error) {
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

	return &MCPServer{client: client}, nil
}

func (s *MCPServer) handleListTools() []Tool {
	return []Tool{
		// Public APIs
		{
			Name:        "get_ticker",
			Description: "Get current ticker information for a trading pair",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "Trading pair (e.g., btc_jpy, eth_jpy)",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "get_trades",
			Description: "Get recent trades for a trading pair",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "Trading pair (e.g., btc_jpy, eth_jpy)",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "get_orderbook",
			Description: "Get order book data for a trading pair",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "Trading pair (e.g., btc_jpy, eth_jpy)",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "get_exchange_status",
			Description: "Get exchange status for a trading pair",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "Trading pair (e.g., btc_jpy, eth_jpy)",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "calculate_order_rate",
			Description: "Calculate order rate for buy/sell orders",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "Trading pair (e.g., btc_jpy)",
						"default":     "btc_jpy",
					},
					"order_type": map[string]interface{}{
						"type":        "string",
						"description": "Order type (buy or sell)",
						"enum":        []string{"buy", "sell"},
					},
					"price": map[string]interface{}{
						"type":        "number",
						"description": "Price per unit",
					},
					"amount": map[string]interface{}{
						"type":        "number",
						"description": "Amount to buy/sell",
					},
				},
				"required": []string{"order_type", "price", "amount"},
			},
		},
		// Account APIs
		{
			Name:        "get_balance",
			Description: "Get account balance information",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "get_accounts",
			Description: "Get account information",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		// Trading APIs
		{
			Name:        "get_transactions",
			Description: "Get transaction history",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "get_open_orders",
			Description: "Get list of open orders",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "create_order",
			Description: "Create a new buy/sell order",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "Trading pair (e.g., btc_jpy)",
						"default":     "btc_jpy",
					},
					"order_type": map[string]interface{}{
						"type":        "string",
						"description": "Order type (buy or sell)",
						"enum":        []string{"buy", "sell"},
					},
					"rate": map[string]interface{}{
						"type":        "number",
						"description": "Order rate (price per unit)",
					},
					"amount": map[string]interface{}{
						"type":        "number",
						"description": "Amount to buy/sell",
					},
				},
				"required": []string{"order_type", "rate", "amount"},
			},
		},
		{
			Name:        "cancel_order",
			Description: "Cancel an existing order",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"order_id": map[string]interface{}{
						"type":        "number",
						"description": "Order ID to cancel",
					},
				},
				"required": []string{"order_id"},
			},
		},
		{
			Name:        "get_order",
			Description: "Get details of a specific order",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
				"properties": map[string]interface{}{
					"order_id": map[string]interface{}{
						"type":        "number",
						"description": "Order ID to retrieve",
					},
				},
				"required": []string{"order_id"},
			},
		},
	}
}

func (s *MCPServer) handleCallTool(toolName string, arguments map[string]interface{}) (interface{}, error) {
	ctx := context.Background()

	switch toolName {
	// Public APIs
	case "get_ticker":
		pair := "btc_jpy"
		if p, ok := arguments["pair"].(string); ok {
			pair = p
		}
		return s.client.REST.Ticker(ctx, pair)

	case "get_trades":
		pair := "btc_jpy"
		if p, ok := arguments["pair"].(string); ok {
			pair = p
		}
		return s.client.REST.Trades(ctx, pair)

	case "get_orderbook":
		pair := "btc_jpy"
		if p, ok := arguments["pair"].(string); ok {
			pair = p
		}
		return s.client.REST.OrderBooks(ctx, pair)

	case "get_exchange_status":
		pair := "btc_jpy"
		if p, ok := arguments["pair"].(string); ok {
			pair = p
		}
		return s.client.REST.ExchangeStatus(ctx, pair)

	case "calculate_order_rate":
		pair := "btc_jpy"
		if p, ok := arguments["pair"].(string); ok {
			pair = p
		}
		orderType := ""
		if ot, ok := arguments["order_type"].(string); ok {
			orderType = ot
		}
		price := 0.0
		if pr, ok := arguments["price"].(float64); ok {
			price = pr
		}
		amount := 0.0
		if am, ok := arguments["amount"].(float64); ok {
			amount = am
		}
		return s.client.REST.OrdersRate(ctx, pair, orderType, price, amount)

	// Account APIs
	case "get_balance":
		return s.client.REST.Balance(ctx)

	case "get_accounts":
		return s.client.REST.Accounts(ctx)

	// Trading APIs
	case "get_transactions":
		return s.client.REST.Transactions(ctx)

	case "get_open_orders":
		return s.client.REST.OpenOrders(ctx)

	case "create_order":
		pair := "btc_jpy"
		if p, ok := arguments["pair"].(string); ok {
			pair = p
		}
		orderType := ""
		if ot, ok := arguments["order_type"].(string); ok {
			orderType = ot
		}
		rate := 0.0
		if r, ok := arguments["rate"].(float64); ok {
			rate = r
		}
		amount := 0.0
		if a, ok := arguments["amount"].(float64); ok {
			amount = a
		}
		return s.client.REST.CreateOrder(ctx, pair, orderType, rate, amount)

	case "cancel_order":
		orderID := 0
		if id, ok := arguments["order_id"].(float64); ok {
			orderID = int(id)
		}
		return s.client.REST.CancelOrder(ctx, orderID)

	case "get_order":
		orderID := 0
		if id, ok := arguments["order_id"].(float64); ok {
			orderID = int(id)
		}
		return s.client.REST.GetOrder(ctx, orderID)

	default:
		return nil, fmt.Errorf("unknown tool: %s", toolName)
	}
}

func (s *MCPServer) handleRequest(req MCPRequest) MCPResponse {
	switch req.Method {
	case "initialize":
		return MCPResponse{
			Jsonrpc: "2.0",
			Result: map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"capabilities": map[string]interface{}{
					"tools": map[string]interface{}{},
				},
				"serverInfo": map[string]interface{}{
					"name":    "coincheck-mcp-server",
					"version": "1.0.0",
				},
			},
			ID: req.ID,
		}
		
	case "tools/list":
		tools := s.handleListTools()
		return MCPResponse{
			Jsonrpc: "2.0",
			Result:  map[string]interface{}{"tools": tools},
			ID:      req.ID,
		}

	case "tools/call":
		toolName, ok := req.Params["name"].(string)
		if !ok {
			return MCPResponse{
				Jsonrpc: "2.0",
				Error:   &MCPError{Code: -32602, Message: "Missing tool name"},
				ID:      req.ID,
			}
		}

		arguments, _ := req.Params["arguments"].(map[string]interface{})
		result, err := s.handleCallTool(toolName, arguments)
		if err != nil {
			return MCPResponse{
				Jsonrpc: "2.0",
				Error:   &MCPError{Code: -32603, Message: err.Error()},
				ID:      req.ID,
			}
		}

		// JSONエンコードして構造体の内容を適切に表示
		resultJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return MCPResponse{
				Jsonrpc: "2.0",
				Error:   &MCPError{Code: -32603, Message: fmt.Sprintf("Failed to marshal result: %v", err)},
				ID:      req.ID,
			}
		}
		
		return MCPResponse{
			Jsonrpc: "2.0",
			Result: map[string]interface{}{"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Tool %s executed successfully. Result:\n%s", toolName, string(resultJSON)),
				},
			}},
			ID: req.ID,
		}

	default:
		return MCPResponse{
			Jsonrpc: "2.0",
			Error:   &MCPError{Code: -32601, Message: "Method not found"},
			ID:      req.ID,
		}
	}
}

func main() {
	server, err := NewMCPServer()
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		var req MCPRequest
		if err := decoder.Decode(&req); err != nil {
			log.Printf("Error decoding request: %v", err)
			continue
		}

		resp := server.handleRequest(req)
		if err := encoder.Encode(resp); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	}
}
