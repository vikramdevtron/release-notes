//+build wireinject

package main

import (
	"github.com/devtron-labs/release-notes/api"
	"github.com/devtron-labs/release-notes/internal/logger"
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		NewApp,
		api.NewMuxRouter,
		logger.NewSugardLogger,
		//logger.NewHttpClient,
		api.NewRestHandlerImpl,
		wire.Bind(new(api.RestHandler), new(*api.RestHandlerImpl)),
	)
	return &App{}, nil
}
