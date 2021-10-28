package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Client *redis.Client

// Init 初始化redis
func Init() (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭redis连接
func Clone()  {
	_ = Client.Close()
}
