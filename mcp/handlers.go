package mcp

import (
	"context"
	"encoding/json"
	"fmt"
)

// HandleRequest processes incoming MCP requests and returns appropriate responses
func (s *Server) HandleRequest(req Request) Response {
	switch req.Method {
	case "notifications/initialized":
		// Notification - no response required
		return Response{Jsonrpc: "", ID: nil}
		
	case "initialize":
		return Response{
			Jsonrpc: "2.0",
			Result: map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"capabilities": map[string]interface{}{
					"tools": map[string]interface{}{},
				},
				"serverInfo": map[string]interface{}{
					"name":    "coincheck-mcp",
					"version": "1.0.0",
				},
			},
			ID: req.ID,
		}

	case "tools/list":
		tools := GetToolDefinitions()
		return Response{
			Jsonrpc: "2.0",
			Result:  map[string]interface{}{"tools": tools},
			ID:      req.ID,
		}

	case "tools/call":
		toolName, ok := req.Params["name"].(string)
		if !ok {
			return Response{
				Jsonrpc: "2.0",
				Error:   &Error{Code: -32602, Message: "Missing tool name"},
				ID:      req.ID,
			}
		}

		arguments, _ := req.Params["arguments"].(map[string]interface{})
		result, err := s.HandleCallTool(toolName, arguments)
		if err != nil {
			return Response{
				Jsonrpc: "2.0",
				Error:   &Error{Code: -32603, Message: err.Error()},
				ID:      req.ID,
			}
		}

		// JSONエンコードして構造体の内容を適切に表示
		resultJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return Response{
				Jsonrpc: "2.0",
				Error:   &Error{Code: -32603, Message: fmt.Sprintf("Failed to marshal result: %v", err)},
				ID:      req.ID,
			}
		}

		return Response{
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
		return Response{
			Jsonrpc: "2.0",
			Error:   &Error{Code: -32601, Message: "Method not found"},
			ID:      req.ID,
		}
	}
}

// HandleCallTool executes the specified tool with given arguments
func (s *Server) HandleCallTool(toolName string, arguments map[string]interface{}) (interface{}, error) {
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

	case "get_order_book":
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

	case "get_account_info":
		return s.client.REST.Accounts(ctx)

	// Trading APIs
	case "get_transaction_history":
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

	case "get_order_details":
		orderID := 0
		if id, ok := arguments["order_id"].(float64); ok {
			orderID = int(id)
		}
		return s.client.REST.GetOrder(ctx, orderID)

	default:
		return nil, fmt.Errorf("unknown tool: %s", toolName)
	}
}