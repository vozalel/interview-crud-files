package middleware

import (
	"fmt"
	"github.com/VictoriaMetrics/metrics"
	"github.com/gin-gonic/gin"
	"time"
)

func MetricsMiddleware(c *gin.Context) {
	startTime := time.Now()
	c.Next()

	metrics.
		GetOrCreateCounter(metricNameWithLables("requests_total", c)).
		Inc()
	metrics.
		GetOrCreateSummary(metricNameWithLables("requests_duration_seconds", c)).
		UpdateDuration(startTime)
}

func metricNameWithLables(name string, c *gin.Context) string {
	return fmt.Sprintf(
		`%s{code="%d",method="%s",path="%s"}`,
		name,
		c.Writer.Status(),
		c.Request.Method,
		c.Request.URL.Path,
	)
}
