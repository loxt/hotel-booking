package main

import (
	"context"
	"github.com/loxt/hotel-booking/booking/pkg/config"
	"github.com/loxt/hotel-booking/booking/pkg/postgres"
	"log/slog"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.NewConfig()
	err := cfg.Load()

	if err != nil {
		logger.Error("failed to load app config", "error", err.Error())

		os.Exit(1)
	}

	_, err = postgres.Connect(ctx, postgres.BuildDSN(cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseName, cfg.DatabasePort))
	if err != nil {
		logger.Error("failed to connect to postgres", "error", err.Error())

		os.Exit(1)
	}

}
