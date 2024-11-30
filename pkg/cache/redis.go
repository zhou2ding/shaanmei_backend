package cache

import (
	"CMAIOT/internal/pkg/conf"
	"CMAIOT/internal/pkg/logger"
	"github.com/go-redis/redis"
)

var gRedis *redis.Client

const (
	BlackListKey   = "TokenBlackList:"
	PackageInfoKey = "PackageInfo:"
)

func InitRedis() error {
	host := conf.Conf.GetString("redis.host")
	pwd := conf.Conf.GetString("redis.password")
	num := conf.Conf.GetInt("redis.num")

	cli := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pwd,
		DB:       num,
	})
	if _, err := cli.Ping().Result(); err != nil {
		logger.Logger.Errorf("connect to redis error: %v", err)
		return err
	}

	gRedis = cli
	return nil
}

func GetRedis() *redis.Client {
	return gRedis
}
