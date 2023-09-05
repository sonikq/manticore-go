package main

import (
	"context"
	cache2 "gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/cache"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/repository"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/services"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/configs"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/db"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
)

func main() {
	// Configurations
	config := configs.Load()

	// Logger
	log := logger.New(config.LogLevel, "api-gateway")
	defer func() {
		err := logger.CleanUp(log)
		log.Error("failed to cleanup logs", logger.Error(err))
	}()

	cache := cache2.New()

	// Initializing app services
	_db, err := db.NewDB(config)
	if err != nil {
		log.Error("failed to initialize db", logger.Error(err))
	}

	repo := repository.NewRepository(cache, _db)

	serviceManager := services.NewService(repo)

	options := api.Option{
		Conf:           config,
		Logger:         log,
		ServiceManager: *serviceManager,
		Cache:          cache,
	}

	server := NewServer(config, api.New(options))

	go func() {
		if err = server.Run(); err != nil {
			log.Fatal("failed to run http server", logger.Error(err))
			panic(err)
		}
	}()

	log.Info("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = server.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", logger.Error(err))
	}

	defer func() {
		if err = _db.Close(); err != nil {
			log.Fatal("failed to close database", logger.Error(err))
		}
	}()
}
