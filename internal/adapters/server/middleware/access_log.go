package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// AccessLog is a middleware function that logs access details of incoming requests.
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time of the request
		startTime := time.Now()

		// Process the request
		c.Next()
		// Request processing is done

		latency := time.Since(startTime)

		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		statusCode := c.Writer.Status()
		responseSize := c.Writer.Size()

		logrus.WithFields(logrus.Fields{
			"method":        method,
			"path":          path,
			"client_ip":     clientIP,
			"status_code":   statusCode,
			"response_size": responseSize,
			"latency":       latency,
		}).Info("Access log")
	}
}
