package main

import (
	"flag"
	"log"

	"github.com/sg3t41/coincheck-mcp/config"
	"github.com/sg3t41/coincheck-mcp/mcp"
)

func main() {
	// フラグ定義
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	var server *mcp.Server
	var err error

	if *configPath != "" {
		// 設定ファイルから読み込み
		cfg, err := config.LoadConfig(*configPath)
		if err != nil {
			log.Fatal(err)
		}
		server, err = mcp.NewServerWithConfig(cfg)
	} else {
		// 環境変数から読み込み（後方互換性）
		server, err = mcp.NewServer()
	}

	if err != nil {
		log.Fatal(err)
	}

	// サーバーを開始
	server.Run()
}

