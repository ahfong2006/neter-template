//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"

	"github.com/Xwudao/neter/internal/cmd"
	"github.com/Xwudao/neter/internal/routes"
	"github.com/Xwudao/neter/pkg/config"
	"github.com/Xwudao/neter/pkg/logger"
)

// wireApp init the application.
func wireApp() (*cobra.Command, func(), error) {
	panic(wire.Build(cmd.NewApp, config.NewConfig, logger.NewLogger, routes.ProviderRouteSet))
}
