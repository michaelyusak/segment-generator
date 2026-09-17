package middleware

import (
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		statusCode := c.Writer.Status()

		entry := logrus.WithFields(logrus.Fields{
			"latency":     time.Since(start),
			"method":      c.Request.Method,
			"path":        path,
			"status_code": statusCode,
		})

		if statusCode >= 400 && statusCode <= 499 {
			for _, err := range c.Errors {
				entry.WithFields(logrus.Fields{
					"error": err,
				}).Warn()
			}
		} else if statusCode >= 500 && statusCode <= 599 {
			for _, err := range c.Errors {
				entry.WithFields(logrus.Fields{
					"error": err,
					"stack": string(debug.Stack()),
				}).Error()
			}
			return
		}
		entry.Info()
	}
}
