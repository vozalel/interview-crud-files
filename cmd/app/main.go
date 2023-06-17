package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/vozalel/interview-crud-files/config"
	"github.com/vozalel/interview-crud-files/internal/app"
	"github.com/vozalel/interview-crud-files/pkg/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("config.NewConfig() err: %v", err))
	}

	logger.Init(cfg.App.Name, cfg.App.Environment, cfg.Logger.Level)

	if cfg.FeatureFlag.DumpConfig {
		spew.Dump(cfg)
	}

	app.Run(cfg)
}
