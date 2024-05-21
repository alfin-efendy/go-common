package database

import (
	"context"

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
	ElasticClient *elasticsearch.Client
	MongoClient   *mongo.Client
)

type Database interface {
	SqlConnection(ctx context.Context, id string) *gorm.DB
	redisConnection(ctx context.Context, id string) *redis.Client
}

type database struct {
	sqlManager   map[string]*gorm.DB
	redisManager map[string]*redis.Client
}
