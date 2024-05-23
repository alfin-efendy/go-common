package app

import (
	"context"

	"github.com/alfin87aa/go-common/configs"
	"github.com/alfin87aa/go-common/database"
	"github.com/alfin87aa/go-common/logger"
	"github.com/alfin87aa/go-common/otel"
	"github.com/alfin87aa/go-common/servers/restapi"
)

func init() {
	ctx := context.Background()
	configs.Load()
	logger.NewLogger().Init()
	otel.Init(ctx)
	database.Init(ctx)
}

func Start() {
	restapi.Run()
}
