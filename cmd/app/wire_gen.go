// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Xwudao/neter/internal/cmd"
	"github.com/Xwudao/neter/internal/routes"
	"github.com/Xwudao/neter/internal/routes/v1"
	"github.com/Xwudao/neter/pkg/config"
	"github.com/Xwudao/neter/pkg/logger"
	"github.com/spf13/cobra"
)

// Injectors from wire.go:

// wireApp init the application.
func wireApp() (*cobra.Command, func(), error) {
	engine := routes.NewEngine()
	koanf, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	sugaredLogger, err := logger.NewLogger(koanf)
	if err != nil {
		return nil, nil, err
	}
	homeRoute := v1.NewHomeRoute(engine, koanf)
	appRoutes := routes.NewAppRoutes(sugaredLogger, homeRoute)
	command := cmd.NewApp(engine, sugaredLogger, appRoutes, koanf)
	return command, func() {
	}, nil
}
