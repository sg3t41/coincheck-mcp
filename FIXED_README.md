# Coincheck MCP Server

CoincheckのAPIを使用したMCP（Model Context Protocol）サーバーです。

## 問題の修正内容

### スキーマ検証エラーの原因と修正

元のエラーログに出ていた `invalid_union`, `invalid_type`, `unrecognized_keys` などのZodErrorは以下の問題が原因でした：

1. **設定ファイルのパス間違い**：
   - 誤: `/home/sg3t41/workspace/auto_investment/coincheck-mcp-server`
   - 正: `/home/sg3t41/workspace/coincheck-mcp-server/coincheck-mcp-server`

2. **JSON Schemaの不完全な定義**：
   - 各ツールのInputSchemaに `$schema` フィールドが欠けていた
   - MCPの標準仕様では `$schema: "http://json-schema.org/draft-07/schema#"` が必要

### 修正箇所

1. **claude_desktop_config.json**:
   ```json
   {
     "mcpServers": {
       "coincheck": {
         "command": "/home/sg3t41/workspace/coincheck-mcp-server/coincheck-mcp-server",
         "env": {
           "COINCHECK_API_KEY": "your_api_key_here",
           "COINCHECK_API_SECRET": "your_api_secret_here"
         }
       }
     }
   }
   ```

2. **main.go**の各ツールのInputSchemaに `$schema` フィールドを追加：
   ```go
   InputSchema: map[string]interface{}{
       "$schema": "http://json-schema.org/draft-07/schema#",
       "type": "object",
       "properties": map[string]interface{}{
           // プロパティ定義
       },
   },
   ```

## 使用方法

1. 環境変数を設定：
   ```bash
   export COINCHECK_API_KEY="your_api_key"
   export COINCHECK_API_SECRET="your_api_secret"
   ```

2. Claude Desktopの設定ファイルにMCPサーバーを追加

3. Claude Desktopを再起動

## 利用可能なツール

- `get_ticker`: 取引ペアの現在価格情報を取得
- `get_trades`: 最近の取引履歴を取得
- `get_orderbook`: オーダーブック情報を取得
- `get_exchange_status`: 取引所のステータスを取得
- `calculate_order_rate`: 注文レートを計算
- `get_balance`: アカウント残高を取得
- `get_accounts`: アカウント情報を取得
- `get_transactions`: 取引履歴を取得
- `get_open_orders`: 未約定注文を取得
- `create_order`: 新しい注文を作成
- `cancel_order`: 注文をキャンセル
- `get_order`: 特定の注文詳細を取得

これで、Claude Desktopを開いた時に出ていたスキーマ検証エラーは解消されるはずです。
