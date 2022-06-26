package data

import (
	"context"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/conf"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/data/ent"
	"github.com/csyangpeng/go-kratos-admin/app/user/service/internal/data/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedisCmd, NewUserRepo)

// Data .
type Data struct {
	db    *ent.Client
	redis redis.Cmdable
}

func NewDB(conf *conf.Data, logger log.Logger) *ent.Client {
	lg := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))

	cli, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		lg.Fatalf("failed opening connection to db: %v", err)
	}

	if err = cli.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		lg.Fatalf("failed create schema resources: %v", err)
	}
	return cli
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	lg := log.NewHelper(log.With(logger, "module", "user-service/data/redis"))

	cli := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := cli.Ping(timeout).Err()
	if err != nil {
		lg.Fatalf("failed connect error: %v", err)
	}
	return cli
}

// NewData .
func NewData(entClient *ent.Client, redisCmd redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	lg := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		db:    entClient,
		redis: redisCmd,
	}

	return d, func() {
		if err := d.db.Close(); err != nil {
			lg.Error(err)
		}
	}, nil
}
