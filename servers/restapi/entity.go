package restapi

import (
	"github.com/alfin87aa/go-common/configs"
	"github.com/etherlabsio/healthcheck/v2"
	"github.com/gin-gonic/gin"
)

var (
	config  = configs.Configs
	Server  *gin.Engine
	options []healthcheck.Option
)
