package global

import (
	"fmt"
	"scaffold-gin/common/config"

	"github.com/go-redis/redis"
)


var REDIS = InitRedis()

func InitRedis() *redis.Client {
	host := config.Conf.REDIS.Host
	port := config.Conf.REDIS.Port
	pass := config.Conf.REDIS.Pass
	dbname := config.Conf.REDIS.DbName

	dsn := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: pass,   // no password set
		DB:       dbname, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		ZAPLOGGER.Sugar().Errorf("redis Init Error %s" , err )
	}

	return client
}
