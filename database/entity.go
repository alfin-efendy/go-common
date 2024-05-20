package database

import (
	"github.com/sirupsen/logrus"
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
)

type Database interface {
}

type database struct {
	sqlManager map[string]*gorm.DB
}
