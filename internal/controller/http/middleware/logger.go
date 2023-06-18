package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vozalel/interview-crud-files/pkg/logger"
)

func LoggerMiddleware(c *gin.Context) {
	c.Next()

	l := logger.
		Instance.
		WithContext(c.Request.Context()).
		WithField("method", c.Request.Method).
		WithField("path", c.Request.URL.Path).
		WithField("ip", c.ClientIP()).
		WithField("proto", c.Request.Proto).
		WithField("userAgent", c.Request.UserAgent()).
		WithField("statusCode", c.Writer.Status()).
		WithField("responseSize", c.Writer.Size())

	for _, err := range c.Errors.Errors() {
		l.Error(err)
	}

	username, _, ok := c.Request.BasicAuth()
	if ok {
		l = l.WithField("username", username)
		l.Info("authorized actions")
	}
}
