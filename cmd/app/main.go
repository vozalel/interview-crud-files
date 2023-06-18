package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/vozalel/interview-crud-files/config"
	"github.com/vozalel/interview-crud-files/internal/app"
	"github.com/vozalel/interview-crud-files/pkg/feature_flag"
	"github.com/vozalel/interview-crud-files/pkg/logger"
	"github.com/vozalel/interview-crud-files/pkg/tracer"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("config.NewConfig() err: %v", err))
	}

	logger.Init(cfg.App.Name, cfg.App.Environment, cfg.Logger.Level)
	feature_flag.Init(cfg.FeatureFlag.DumpConfig, cfg.FeatureFlag.TraceEnabled)

	if feature_flag.Get().DumpConfigEnabled() {
		spew.Dump(cfg)
	}

	if feature_flag.Get().TraceEnabled() {
		err = tracer.Init(cfg.Trace.URL, cfg.App.Name, cfg.App.Environment)
		if err != nil {
			logger.Instance.Fatal(fmt.Errorf("main - tracer.Init: %w", err))
		}
	}

	app.Run(cfg)
}
