package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TiJon8/shorter-go/internal/config"
	"github.com/TiJon8/shorter-go/internal/features"
	log "github.com/TiJon8/shorter-go/internal/logger"
	"github.com/TiJon8/shorter-go/internal/server"
	"github.com/TiJon8/shorter-go/internal/storage"
)



func main() {
	appCfg := config.AppConfigMust()
	fmt.Println(appCfg)
	logger := log.NewLogger(appCfg.Env, appCfg.LogLevel)

	storage, err := storage.Init(appCfg.StoragePath)
	if err != nil {
		logger.Error("Open connection to databse failed", log.ErrorAttr("error", err))
		os.Exit(1)
	}
	logger.Info(("Sqlite has oppened"))

	router := features.InitRouter(storage, logger)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	serverCfg := config.ServerConfigMust()
	Server := server.NewHTTPServer(serverCfg, logger.Logger, router, storage)
	if err := Server.Run(ctx); err != nil {
		logger.Error("Server error", log.ErrorAttr("err", err))
	}
}