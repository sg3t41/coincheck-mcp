# Coincheck MCP Server

Coincheck暗号通貨取引所APIと統合するModel Context Protocol (MCP)サーバー。Claude DesktopからCoincheckサービスを操作できます。

## 機能

### パブリックAPI
- **get_ticker** - 取引ペアの現在のティッカー情報を取得
- **get_trades** - 取引ペアの最近の取引履歴を取得
- **get_orderbook** - 板情報を取得
- **get_exchange_status** - 取引所のステータスを確認
- **calculate_order_rate** - 売買注文のレートを計算

### アカウントAPI
- **get_balance** - アカウントの残高情報を取得
- **get_accounts** - アカウント情報を取得

### 取引API
- **get_transactions** - 取引履歴を取得
- **get_open_orders** - 未約定の注文一覧を取得
- **create_order** - 新規売買注文を作成
- **cancel_order** - 既存の注文をキャンセル
- **get_order** - 特定の注文の詳細を取得

## 必要条件

- Go 1.21以上
- Coincheck APIクレデンシャル（APIキーとシークレット）
- Claude Desktop

## インストール

1. リポジトリをクローン:
```bash
git clone https://github.com/sg3t41/coincheck-mcp-server.git
cd coincheck-mcp-server
```

2. サーバーをビルド:
```bash
go build -o coincheck-mcp-server main.go
```

## 設定

### 1. APIクレデンシャルの設定

[Coincheck](https://coincheck.com/ja/exchange/api_settings)からAPIクレデンシャルを取得してください。

### 2. Claude Desktopの設定

Claude Desktopの設定ファイルに以下の設定を追加してください：

**設定ファイルの場所:**
- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`
- Linux: `~/.config/Claude/claude_desktop_config.json`

**設定内容:**
```json
{
  "mcpServers": {
    "coincheck": {
      "command": "/path/to/coincheck-mcp-server",
      "env": {
        "COINCHECK_API_KEY": "your_api_key_here",
        "COINCHECK_API_SECRET": "your_api_secret_here"
      }
    }
  }
}
```

以下を置き換えてください:
- `/path/to/coincheck-mcp-server` - ビルドしたバイナリへの実際のパス
- `your_api_key_here` - あなたのCoincheck APIキー
- `your_api_secret_here` - あなたのCoincheck APIシークレット

### 3. 環境変数の使用（オプション）

設定ファイルにクレデンシャルをハードコーディングする代わりに、シェルスクリプトを作成できます：

```bash
#!/bin/bash
export COINCHECK_API_KEY="${COINCHECK_API_KEY}"
export COINCHECK_API_SECRET="${COINCHECK_API_SECRET}"
exec /path/to/coincheck-mcp-server
```

その後、Claude Desktopの設定をスクリプトを使用するように更新：
```json
{
  "mcpServers": {
    "coincheck": {
      "command": "/path/to/run-coincheck-mcp.sh"
    }
  }
}
```

## 使用方法

1. Claude Desktopを起動
2. `@coincheck`と入力して利用可能なツールを確認
3. 自然言語でCoincheckと対話：

### 使用例
```
@coincheck BTCの現在価格を取得
@coincheck アカウントの残高を表示
@coincheck DOGEの板情報を取得
@coincheck 100 DOGEを25円で買い注文を作成
@coincheck 注文12345をキャンセル
```

## サポートされている取引ペア

- btc_jpy（ビットコイン）
- eth_jpy（イーサリアム）
- etc_jpy（イーサリアムクラシック）
- lsk_jpy（リスク）
- mona_jpy（モナコイン）
- plt_jpy（パレットトークン）
- fnct_jpy（フィネクサス）
- dai_jpy（ダイ）
- wbtc_jpy（ラップドビットコイン）
- doge_jpy（ドージコイン）
- その他多数...

## セキュリティに関する注意事項

- APIクレデンシャルを絶対にバージョン管理にコミットしない
- APIキーを安全に保管し、定期的にローテーションする
- 可能な限り環境変数でクレデンシャルを管理する
- 注文作成時は注意深く - 金額とレートを必ず二重確認する

## トラブルシューティング

### MCPサーバーが接続されない場合
1. Claude Desktop設定のバイナリパスが正しいか確認
2. バイナリに実行権限があるか確認: `chmod +x coincheck-mcp-server`
3. APIクレデンシャルが正しく設定されているか確認
4. ログを確認: `~/.config/Claude/logs/mcp-server-coincheck.log`

### APIエラー
- APIキーに必要な権限があるか確認
- 有効な取引ペアを使用しているか確認
- 取引に必要な残高があるか確認

## 開発

### ソースからのビルド
```bash
go mod download
go build -o coincheck-mcp-server main.go
```


## ライセンス

MITライセンス

## 謝辞

- [go-coincheck](https://github.com/sg3t41/go-coincheck)ライブラリを使用して構築
- [Model Context Protocol](https://modelcontextprotocol.io/)を実装
- [Claude Desktop](https://claude.ai/download)で使用