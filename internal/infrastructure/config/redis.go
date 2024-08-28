package config

import (
	"time"

	"github.com/go-redis/redis/v8"
)

func ConnMasterRedis(env_db_uri string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:         env_db_uri,
		Password:     "",
		DB:           0,               //Num of DB, 0 by default
		PoolSize:     80,              //Size of the pool
		MinIdleConns: 20,              //Min conn in the pool
		DialTimeout:  5 * time.Second, //Timeout initial conns
		ReadTimeout:  3 * time.Second, //Timeout read
		WriteTimeout: 3 * time.Second, //Timeout write
	})

	return rdb
}
