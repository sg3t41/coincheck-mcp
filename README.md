# Coincheck MCP Server

A Model Context Protocol (MCP) server that integrates with the Coincheck cryptocurrency exchange API, allowing Claude Desktop to interact with Coincheck services.

## Features

### Public APIs
- **get_ticker** - Get current ticker information for trading pairs
- **get_trades** - Get recent trades for a trading pair
- **get_orderbook** - Get order book data
- **get_exchange_status** - Check exchange status
- **calculate_order_rate** - Calculate order rates for buy/sell orders

### Account APIs
- **get_balance** - Get account balance information
- **get_accounts** - Get account information

### Trading APIs
- **get_transactions** - Get transaction history
- **get_open_orders** - Get list of open orders
- **create_order** - Create new buy/sell orders
- **cancel_order** - Cancel existing orders
- **get_order** - Get details of specific orders

## Requirements

- Go 1.21 or higher
- Coincheck API credentials (API Key and Secret)
- Claude Desktop

## Installation

1. Clone the repository:
```bash
git clone https://github.com/sg3t41/coincheck-mcp-server.git
cd coincheck-mcp-server
```

2. Build the server:
```bash
go build -o coincheck-mcp-server main.go
```

## Configuration

### 1. Set up API Credentials

Get your API credentials from [Coincheck](https://coincheck.com/ja/exchange/api_settings).

### 2. Configure Claude Desktop

Add the following configuration to your Claude Desktop config file:

**Location:**
- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`
- Linux: `~/.config/Claude/claude_desktop_config.json`

**Configuration:**
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

Replace:
- `/path/to/coincheck-mcp-server` with the actual path to your built binary
- `your_api_key_here` with your Coincheck API key
- `your_api_secret_here` with your Coincheck API secret

### 3. Using Environment Variables (Optional)

Instead of hardcoding credentials in the config, you can create a shell script:

```bash
#!/bin/bash
export COINCHECK_API_KEY="${COINCHECK_API_KEY}"
export COINCHECK_API_SECRET="${COINCHECK_API_SECRET}"
exec /path/to/coincheck-mcp-server
```

Then update your Claude Desktop config to use the script:
```json
{
  "mcpServers": {
    "coincheck": {
      "command": "/path/to/run-coincheck-mcp.sh"
    }
  }
}
```

## Usage

1. Start Claude Desktop
2. Type `@coincheck` to see available tools
3. Use natural language to interact with Coincheck:

### Examples
```
@coincheck get the current BTC price
@coincheck show my account balance
@coincheck get order book for DOGE
@coincheck create a buy order for 100 DOGE at 25 JPY
@coincheck cancel order 12345
```

## Supported Trading Pairs

- btc_jpy (Bitcoin)
- eth_jpy (Ethereum)
- etc_jpy (Ethereum Classic)
- lsk_jpy (Lisk)
- mona_jpy (Monacoin)
- plt_jpy (Palette Token)
- fnct_jpy (Finnexus)
- dai_jpy (Dai)
- wbtc_jpy (Wrapped Bitcoin)
- doge_jpy (Dogecoin)
- And more...

## Security Notes

- Never commit your API credentials to version control
- Keep your API keys secure and rotate them regularly
- Use environment variables for credentials when possible
- Be cautious when creating orders - always double-check amounts and rates

## Troubleshooting

### MCP Server Not Connecting
1. Check that the binary path in Claude Desktop config is correct
2. Ensure the binary has execute permissions: `chmod +x coincheck-mcp-server`
3. Verify API credentials are correctly set
4. Check logs at `~/.config/Claude/logs/mcp-server-coincheck.log`

### API Errors
- Ensure your API key has the necessary permissions
- Check that you're using valid trading pairs
- Verify your account has sufficient balance for trades

## Development

### Building from Source
```bash
go mod download
go build -o coincheck-mcp-server main.go
```

### Running Tests
```bash
go test ./...
```

## License

MIT License

## Acknowledgments

- Built with [go-coincheck](https://github.com/sg3t41/go-coincheck) library
- Implements the [Model Context Protocol](https://modelcontextprotocol.io/)
- For use with [Claude Desktop](https://claude.ai/download)