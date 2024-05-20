package restapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/alfin87aa/go-common/configs"
	"github.com/alfin87aa/go-common/logger"
	"github.com/etherlabsio/healthcheck/v2"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func init() {
	config = configs.Configs
	Server = gin.Default()

	Server.Use(
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
		otelgin.Middleware(config.App.Name),
		RequestIDMiddleware(),
		LoggerMiddleware(),
		CORSMiddleware(),
		HelmetMiddleware(),
	)

	Server.GET("/health", gin.WrapH(healthz()))
}

func AddChecker(name string, f func(ctx context.Context) error) {
	options = append(
		options,
		healthcheck.WithChecker(
			name,
			healthcheck.CheckerFunc(f),
		))
}

// @Summary		Health Check
// @Description	Perform a health check
// @Produce		json
// @Success		200
// @Failure		503
// @Router			/healthz [get]
func healthz() http.Handler {
	options = append(options, healthcheck.WithTimeout(5*time.Second))
	return healthcheck.Handler(options...)
}

func Run() {
	ctx := context.Background()

	if config.Server.RestAPI == nil || !config.Server.RestAPI.Enable {
		logger.Warn(ctx, "REST API is disabled")
		return
	}

	port := config.Server.RestAPI.Port

	err := Server.Run(fmt.Sprintf("%s:%d", "0.0.0.0", port))
	if err != nil {
		logger.Fatalf(ctx, err, "Failed to run REST server, port=%d", port)
	}
}
