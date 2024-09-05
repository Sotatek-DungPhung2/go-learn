package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ANSI Color Codes
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// LoggingMiddleware is a middleware that logs the requests and responses
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate response time
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get status code and method
		statusCode := c.Writer.Status()
		method := c.Request.Method

		// Set color based on status code
		var color string
		switch {
		case statusCode >= 200 && statusCode < 300:
			color = Green
		case statusCode >= 300 && statusCode < 400:
			color = Blue
		case statusCode >= 400 && statusCode < 500:
			color = Yellow
		case statusCode >= 500:
			color = Red
		default:
			color = White
		}

		// Log the request with colored output
		fmt.Printf("%s[%s] %s %s %s | %d | %s in %v%s\n",
			color, time.Now().Format("2006-01-02 15:04:05"),
			method,
			c.Request.URL.Path,
			c.ClientIP(),
			statusCode,
			httpStatusText(statusCode),
			latency,
			Reset,
		)
	}
}

// httpStatusText returns the status text based on status code
func httpStatusText(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
}
