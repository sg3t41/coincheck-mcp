package mcp

// GetToolDefinitions returns all available tool definitions for the Coincheck MCP server
func GetToolDefinitions() []Tool {
	return []Tool{
		// マーケット情報系
		{
			Name:        "get_ticker",
			Description: "指定通貨ペアの現在価格情報を取得",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "取引ペア（例：btc_jpy, eth_jpy）",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "get_trades",
			Description: "指定通貨ペアの最近の取引履歴を取得",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "取引ペア（例：btc_jpy, eth_jpy）",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "get_order_book",
			Description: "指定通貨ペアの売買注文一覧（板情報）を取得",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "取引ペア（例：btc_jpy, eth_jpy）",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "get_exchange_status",
			Description: "指定通貨ペアの取引所稼働状態を取得",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "取引ペア（例：btc_jpy, eth_jpy）",
						"default":     "btc_jpy",
					},
				},
			},
		},
		{
			Name:        "calculate_order_rate",
			Description: "売買注文のレートを計算",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "取引ペア（例：btc_jpy）",
						"default":     "btc_jpy",
					},
					"order_type": map[string]interface{}{
						"type":        "string",
						"description": "注文種別（buy または sell）",
						"enum":        []string{"buy", "sell"},
					},
					"price": map[string]interface{}{
						"type":        "number",
						"description": "単価",
					},
					"amount": map[string]interface{}{
						"type":        "number",
						"description": "数量",
					},
				},
				"required": []string{"order_type", "price", "amount"},
			},
		},
		// アカウント系
		{
			Name:        "get_balance",
			Description: "アカウントの残高情報を取得",
			InputSchema: map[string]interface{}{
				"$schema":    "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "get_account_info",
			Description: "アカウントの詳細情報を取得",
			InputSchema: map[string]interface{}{
				"$schema":    "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		// 取引系
		{
			Name:        "get_transaction_history",
			Description: "アカウントの取引履歴を取得",
			InputSchema: map[string]interface{}{
				"$schema":    "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "get_open_orders",
			Description: "現在の未約定注文一覧を取得",
			InputSchema: map[string]interface{}{
				"$schema":    "http://json-schema.org/draft-07/schema#",
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "create_order",
			Description: "新しい売買注文を作成",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"pair": map[string]interface{}{
						"type":        "string",
						"description": "取引ペア（例：btc_jpy）",
						"default":     "btc_jpy",
					},
					"order_type": map[string]interface{}{
						"type":        "string",
						"description": "注文種別（buy または sell）",
						"enum":        []string{"buy", "sell"},
					},
					"rate": map[string]interface{}{
						"type":        "number",
						"description": "注文レート（単価）",
					},
					"amount": map[string]interface{}{
						"type":        "number",
						"description": "数量",
					},
				},
				"required": []string{"order_type", "rate", "amount"},
			},
		},
		{
			Name:        "cancel_order",
			Description: "既存の注文をキャンセル",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"order_id": map[string]interface{}{
						"type":        "number",
						"description": "キャンセルする注文ID",
					},
				},
				"required": []string{"order_id"},
			},
		},
		{
			Name:        "get_order_details",
			Description: "特定注文の詳細情報を取得",
			InputSchema: map[string]interface{}{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type":    "object",
				"properties": map[string]interface{}{
					"order_id": map[string]interface{}{
						"type":        "number",
						"description": "取得する注文ID",
					},
				},
				"required": []string{"order_id"},
			},
		},
	}
}