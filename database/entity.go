package database

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	PostgreSQL = "postgresql"
	MySQL      = "mysql"
	MSSQL      = "mssql"
)

var (
	sql []string = []string{
		PostgreSQL,
		MySQL,
		MSSQL,
	}
	logrusLevelMap = map[logrus.Level]logger.LogLevel{
		logrus.PanicLevel: logger.Error,
		logrus.FatalLevel: logger.Error,
		logrus.ErrorLevel: logger.Error,
		logrus.WarnLevel:  logger.Warn,
		logrus.InfoLevel:  logger.Warn,
		logrus.DebugLevel: logger.Info,
		logrus.TraceLevel: logger.Info,
	}
	elasticClient *elasticsearch.Client
	mongoClient   *mongo.Client
	sqlManager    map[string]*gorm.DB
	redisManager  map[string]*redis.Client
)
