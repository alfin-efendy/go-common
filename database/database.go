package database

import (
	"context"
	"fmt"
	"time"

	"github.com/alfin87aa/go-common/configs"
	log "github.com/alfin87aa/go-common/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (d *database) Setup(ctx context.Context) {
	config := configs.Configs.DB
	logLevel := log.GetLevel()

	for id, dbConfig := range config {
		var dialector gorm.Dialector

		switch dbConfig.Driver {
		case MySQL:
			// setup MySQL database
		case PostgreSQL:
			// setup PostgreSQL database
		case MSSQL:
			// setup MSSQL database
		default:
			log.Fatalf(ctx, fmt.Errorf("database driver %s is not supported", dbConfig.Driver), "❌ Failed unsupported database driver")
			return
		}

		loggerConfig := logger.Config{
			SlowThreshold:             3 * time.Second,
			LogLevel:                  logrusLevelMap[logLevel],
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		}

		db, err := gorm.Open(dialector, &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.New(log.GetLogrus(), loggerConfig),
		})

		if err != nil {
			log.Fatalf(ctx, err, "❌ Failed to open database connection")
			return
		}

		log.Infof(ctx, "✅ Database connection established, id=%s", id)

		db.Session(&gorm.Session{
			FullSaveAssociations: true,
			PrepareStmt:          true,
		})

		dbSql, err := db.DB()
		if err != nil {
			log.Fatalf(ctx, err, "❌ Failed to get database connection")
			return
		}

		dbSql.SetMaxIdleConns(dbConfig.PoolingConnection.MaxIdle)
		dbSql.SetMaxOpenConns(dbConfig.PoolingConnection.MaxOpen)
		dbSql.SetConnMaxLifetime(time.Duration(dbConfig.PoolingConnection.MaxLifetime) * time.Second)

		d.sqlManager[id] = db
	}
}
