// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cmd"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/routes"
	"github.com/Xwudao/neter-template/internal/routes/v1"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
	"github.com/Xwudao/neter-template/pkg/utils/cron"
)

// Injectors from wire.go:

func mainApp() (*cmd.MainApp, func(), error) {
	koanf, err := config.NewKoanf()
	if err != nil {
		return nil, nil, err
	}
	appConfig, err := config.NewConfig(koanf)
	if err != nil {
		return nil, nil, err
	}
	sugaredLogger, err := logger.NewLogger(koanf)
	if err != nil {
		return nil, nil, err
	}
	engine := routes.NewEngine(appConfig, sugaredLogger)
	appContext := core.NewAppContext()
	userBiz := biz.NewUserBiz()
	userRoute := v1.NewUserRoute(engine, userBiz, koanf)
	httpEngine, err := routes.NewHttpEngine(engine, appConfig, sugaredLogger, appContext, userRoute)
	if err != nil {
		return nil, nil, err
	}
	cronCron := cron.NewCron(sugaredLogger)
	initSystem := core.NewInitSystem(cronCron)
	cmdMainApp, cleanup := cmd.NewMainApp(httpEngine, sugaredLogger, koanf, cronCron, initSystem)
	return cmdMainApp, func() {
		cleanup()
	}, nil
}
