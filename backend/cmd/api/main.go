package main

import (
	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/server"
	"github.com/blastertwist/flag-dash/pkg/database"
	"github.com/blastertwist/flag-dash/pkg/logger"
)

func main(){
	
	// Load Config
	var cfg config.Config
	cfg.LoadConfig("dev-env")

	// Initialize Logger
	var logger = logger.NewLogger(&cfg)
	logger.InitLogger()

	// Initialize DB Connection
	conn, _ := database.ConnectDB(&cfg)
	defer conn.Close()

	// Initialize Server
	server := server.NewServer(&cfg, conn, logger)
	server.Run()
	
}