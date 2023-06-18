package http

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vozalel/interview-crud-files/internal/controller/http/middleware"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/pkg/feature_flag"
	"net/http"

	// Swagger docs.
	//_ "github.com/vozalel/interview-crud-files/docs"
)

// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /
func NewRouter(handler *gin.Engine, datasourceUC entity.IDatasourceUC) {
	// Service endpoints:
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	handler.Any("/metrics", func(c *gin.Context) { metrics.WritePrometheus(c.Writer, true) })

	// Middleware:
	handler.Use(
		gin.CustomRecovery(middleware.RecoveryMiddleware),
		middleware.CORSMiddleware(),
		middleware.MetricsMiddleware,
		middleware.LoggerMiddleware,
	)

	if feature_flag.Get().TraceEnabled() {
		handler.Use(middleware.TracerMiddleware)
	}

	// Business endpoints:
	handlerDatasource := handler.Group("/source")
	{
		newDatasourceRoutes(handlerDatasource, datasourceUC)
	}

	handlerDatasourceList := handler.Group("/list")

	{
		newDatasourceListRoutes(handlerDatasourceList, datasourceUC)
	}
}
