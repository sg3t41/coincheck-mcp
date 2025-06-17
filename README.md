# Coincheck MCP

Coincheck暗号通貨取引所のAPIをClaude Desktopから使えるようにするMCPサーバー。

## 🚀 クイックスタート

```bash
# 1. クローン & ビルド
git clone https://github.com/sg3t41/coincheck-mcp-server.git
cd coincheck-mcp
make all

# 2. 設定
make setup
# config/coincheck_config.json を編集してAPIキーを設定

# 3. Claude Desktopに設定をデプロイ
make deploy
```

## 🔧 利用可能なコマンド

```bash
make help     # ヘルプ表示
make all      # ビルド + インストール（デフォルト）
make setup    # 設定ファイル作成
make test     # テスト実行
make deploy   # Claude Desktopに設定をデプロイ
make status   # プロジェクト状態表示
make clean    # クリーンアップ
```

## ⚙️ 設定

### APIキーの取得
[Coincheck API設定ページ](https://coincheck.com/ja/exchange/api_settings)でAPIキーを取得してください。

### 設定ファイルの編集
```bash
make setup  # config/coincheck_config.json を作成
```

作成された `config/coincheck_config.json` を編集：
```json
{
  "api_key": "your_api_key_here",
  "api_secret": "your_api_secret_here"
}
```

### Claude Desktop設定（自動）
```bash
make deploy  # 自動でClaude Desktopに設定をコピー
```

### Claude Desktop設定（手動）
設定ファイルの場所：
- **Linux**: `~/.config/Claude/claude_desktop_config.json`
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\\Claude\\claude_desktop_config.json`

設定例：
```json
{
  "mcpServers": {
    "coincheck": {
      "command": "coincheck-mcp",
      "args": ["--config", "/path/to/config/coincheck_config.json"]
    }
  }
}
```

## 🎯 使用方法

Claude Desktopで `@coincheck` と入力して利用可能なツールを確認。

### 使用例
```
@coincheck BTCの現在価格を取得
@coincheck アカウントの残高を表示
@coincheck DOGEの板情報を取得
@coincheck 100 DOGEを25円で買い注文を作成
```

## 📊 機能一覧

### パブリックAPI
- 価格取得、取引履歴、板情報、取引所状態、注文レート計算

### アカウントAPI
- 残高確認、アカウント情報取得

### 取引API
- 取引履歴、注文作成・キャンセル・詳細取得、未約定注文一覧

## 🔒 セキュリティ

- APIクレデンシャルはGitにコミットされません
- 機密情報は `.gitignore` で除外済み
- 注文作成時は金額とレートを必ず確認してください

## 🛠️ 開発

```bash
make status   # 現在の状態確認
make clean    # クリーンアップ
make rebuild  # 強制再ビルド
```

## ⚡ トラブルシューティング

**MCPサーバーが接続されない場合:**
```bash
make test  # バイナリテスト
make status  # 状態確認
```

**設定を確認:**
- `coincheck_config.json` にAPIキーが正しく設定されているか
- Claude Desktopを再起動したか
- バイナリに実行権限があるか（`make test`で確認）

## 📄 ライセンス

MIT License

---

**Built with [go-coincheck](https://github.com/sg3t41/go-coincheck) • [MCP](https://modelcontextprotocol.io/)**
