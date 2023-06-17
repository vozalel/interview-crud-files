package app

import (
	"github.com/vozalel/interview-crud-files/config"
	"github.com/vozalel/interview-crud-files/pkg/logger"
	"github.com/vozalel/interview-crud-files/pkg/postgres"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(
		cfg.Postgresql.URL,
		logger.Instance,
	)
	if err != nil {
		logger.Instance.Fatal(err)
	}
	defer pg.Close()

	logger.Instance.Debug("app started...")
}
