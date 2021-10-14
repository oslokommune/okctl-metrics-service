package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func generateJSONLoggerMiddleware(logger *logrus.Logger, skipPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if contains(skipPaths, c.Request.URL.Path) {
			return
		}

		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := time.Since(start)

		entry := logger.WithFields(logrus.Fields{
			"client_ip": c.Request.RemoteAddr,
			"duration":  duration.Microseconds(),
			"method":    c.Request.Method,
			"path":      c.Request.RequestURI,
			"status":    c.Writer.Status(),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
