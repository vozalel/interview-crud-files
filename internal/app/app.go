package app

import (
	"github.com/vozalel/interview-crud-files/cmd/config"
	"github.com/vozalel/interview-crud-files/pkg/logger"
)

func Run(cfg *config.Config) {
	logger.Instance.Debug("app started...")
}
