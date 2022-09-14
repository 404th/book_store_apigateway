package main

import (
	// "fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/404th/book_store_apigateway/api"
	"github.com/404th/book_store_apigateway/config"
	"github.com/404th/book_store_apigateway/pkg/logger"
	"github.com/404th/book_store_apigateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "book_store_apigateway")

	gprcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	quit := make(chan os.Signal, 1)
	go server.Run(cfg.HttpPort)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	log.Info("Server exiting")
}
