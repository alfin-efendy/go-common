package database

import (
	"context"
	"fmt"
	"time"

	"github.com/alfin87aa/go-common/configs"
	log "github.com/alfin87aa/go-common/logger"
	"github.com/alfin87aa/go-common/servers/restapi"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (d *database) initSql(ctx context.Context) {
	config := configs.Configs.DB.Sql
	logLevel := log.GetLevel()

	for id, dbConfig := range config {
		var dialector gorm.Dialector

		switch dbConfig.Driver {
		case MySQL:
			dialector = mysql.Open(
				fmt.Sprintf(
					"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
					dbConfig.Username,
					dbConfig.Password,
					dbConfig.Host,
					dbConfig.Port,
					dbConfig.Database,
				),
			)
		case PostgreSQL:
			dialector = postgres.Open(
				fmt.Sprintf(
					"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
					dbConfig.Host,
					dbConfig.Port,
					dbConfig.Username,
					dbConfig.Database,
					dbConfig.Password,
				),
			)
		case MSSQL:
			dialector = sqlserver.Open(
				fmt.Sprintf(
					"sqlserver://%s:%s@%s:%d?database=%s",
					dbConfig.Username,
					dbConfig.Password,
					dbConfig.Host,
					dbConfig.Port,
					dbConfig.Database,
				),
			)
		default:
			log.Fatalf(ctx, fmt.Errorf("sql database driver %s is not supported", dbConfig.Driver), "❌ Failed unsupported database driver")
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
			log.Fatalf(ctx, err, "❌ Failed to open sql database connection")
			return
		}

		log.Infof(ctx, "✅ Database connection established, id=%s", id)

		db.Session(&gorm.Session{
			FullSaveAssociations: true,
			PrepareStmt:          true,
		})

		dbSql, err := db.DB()
		if err != nil {
			log.Fatalf(ctx, err, "❌ Failed to get sql database connection")
			return
		}

		dbSql.SetMaxIdleConns(dbConfig.PoolingConnection.MaxIdle)
		dbSql.SetMaxOpenConns(dbConfig.PoolingConnection.MaxOpen)
		dbSql.SetConnMaxLifetime(time.Duration(dbConfig.PoolingConnection.MaxLifetime) * time.Second)

		restapi.AddChecker("sql-"+id, func(ctx context.Context) error {
			return dbSql.Ping()
		})

		d.sqlManager[id] = db
	}
}

func (d *database) SqlConnection(ctx context.Context, id string) *gorm.DB {
	if db, ok := d.sqlManager[id]; ok {
		return db
	}

	log.Fatalf(ctx, fmt.Errorf("sql database connection %s not found", id), "❌ Failed to get sql database connection")
	return nil
}
