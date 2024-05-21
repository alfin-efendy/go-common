package database

import (
	"context"
)

func (d *database) Setup(ctx context.Context) {
	go d.initSql(ctx)
	go d.initRedis(ctx)
	go d.initElasticsearch(ctx)
	go d.initMongo(ctx)
}
