package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/TiJon8/shorter-go/pkg/config"
	"github.com/TiJon8/shorter-go/pkg/features"
	log "github.com/TiJon8/shorter-go/pkg/logger"
	"github.com/TiJon8/shorter-go/pkg/server"
	"github.com/TiJon8/shorter-go/pkg/storage"
)

var (
	useEnv = flag.Bool("use-env", false, "if true command will use .env file")
	addr = flag.String("addr", "", "server listening on [:addr]")
	storagePath = flag.String("storage-path", "./database.db", "path for sqlite .db file")
	logLevel = flag.String("log-level", "info", "minimum level for logger")
	env = flag.String("env", "dev", "env mode")
)

func main() {
	flag.Parse()

	appCfg := new(config.AppConfig)
	serverCfg := new(config.ServerConfig)
	if *useEnv {
		appCfg = config.AppConfigMust()
		serverCfg = config.ServerConfigMust()
	} else {
		if *addr == "" {
			flag.Usage()
			return
		}
		serverCfg.Addr = *addr
		appCfg.StoragePath = *storagePath
		appCfg.LogLevel = *logLevel
		appCfg.Env = *env
	}

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
	Server := server.NewHTTPServer(serverCfg, nil, router, storage)
	if err := Server.Run(ctx); err != nil {
		logger.Error("Server error", log.ErrorAttr("err", err))
	}
}