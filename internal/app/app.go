package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vozalel/interview-crud-files/config"
	"github.com/vozalel/interview-crud-files/internal/adapter/acl"
	"github.com/vozalel/interview-crud-files/internal/adapter/datasource"
	"github.com/vozalel/interview-crud-files/internal/controller/http"
	"github.com/vozalel/interview-crud-files/internal/usecase"
	"github.com/vozalel/interview-crud-files/pkg/http_server"
	"github.com/vozalel/interview-crud-files/pkg/logger"
	"github.com/vozalel/interview-crud-files/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
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

	aclManager := acl.New(pg)
	fileManager := datasource.New(cfg.Datasource.Path)

	datasourceUC := usecase.New(aclManager, fileManager)

	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	http.NewRouter(handler, datasourceUC)
	httpServer := http_server.New(handler, http_server.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Instance.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Instance.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		logger.Instance.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
