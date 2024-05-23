package database

import (
	"context"
)

func Init(ctx context.Context) {
	go initSql(ctx)
	go initRedis(ctx)
	go initElasticsearch(ctx)
	go initMongo(ctx)
}
