package restapi

import (
	"strconv"
	"time"

	"github.com/alfin87aa/go-common/configs"
	"github.com/alfin87aa/go-common/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// RequestIDMiddleware generates a unique request ID and sets it in the response header.
// It is a middleware function for Gin framework.
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
		c.Next()
	}
}

// LoggerMiddleware logs the request and response details.
// It is a middleware function for Gin framework.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// Log details
		logger.Info(
			c.Request.Context(),
			"Request",
			logrus.Fields{
				"method":  c.Request.Method,
				"path":    c.Request.URL.Path,
				"status":  c.Writer.Status(),
				"latency": latency.String(),
			},
		)
	}
}

// CORSMiddleware adds CORS headers to the response.
// It is a middleware function for Gin framework.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if configs.Configs.Server.RestAPI.Cors == nil {
			c.Next()
		}

		config := configs.Configs.Server.RestAPI.Cors
		for _, origin := range config.AllowOrigins {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		for _, method := range config.AllowMethods {
			c.Writer.Header().Set("Access-Control-Allow-Methods", method)
		}
		for _, header := range config.AllowHeaders {
			c.Writer.Header().Set("Access-Control-Allow-Headers", header)
		}
		for _, exposeHeaders := range config.ExposeHeaders {
			c.Writer.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
		}
		c.Writer.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.MaxAge))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.AllowCredentials))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// HelmetMiddleware adds security headers to the response.
func HelmetMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY always")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
		c.Writer.Header().Set("Cache-control", "no-store")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		c.Writer.Header().Set("Permissions-Policy", "geolocation=(), midi=(), notifications=(), push=(), sync-xhr=(), microphone=(), camera=(), magnetometer=(), gyroscope=(), speaker=(), vibrate=(), fullscreen=(self), payment=()")
		c.Next()
	}
}
