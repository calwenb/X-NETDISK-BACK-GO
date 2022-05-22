package util

import (
	"github.com/go-redis/redis"
	"netdisk-back/app/config"
)

var rdb *redis.Client
var (
	addr     = config.Get("redis.addr")
	password = config.Get("redis.password")
)

// 初始化连接
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return
	}

}
func GetRDBConn() *redis.Client {
	return rdb
}
