package main

import (
	"log"
	"time"

	"github.com/JuncYoung/libredis"
)

func main() {
	var redisPair libredis.RedisGopher
	rdsDeployKind := "1"
	redisMain, pwd := "127.0.0.1:6379", ""
	switch rdsDeployKind {
	case libredis.REDIS_DEPLOY_KIND_OFFICIAL_CLUSTER:
		log.Printf("redispair.NewRedisGoV6Cluster: %s\n", "redispair.NewRedisGoV6Cluster")
		redisPair = libredis.NewRedisGoV6Cluster(redisMain, pwd)
	case libredis.REDIS_DEPLOY_KIND_SINGLE:
		log.Printf("redispair.NewRedisGoV6: %s\n", "redispair.NewRedisGoV6")
		redisPair = libredis.NewRedisGoV6(redisMain, pwd)
	default:
		redisPair = libredis.NewRedisGoV6(redisMain, pwd)
	}

	err := redisPair.Set("myKey", 66.666, 100*time.Second).Err()
	if err != nil {
		panic(err.Error())
	}

	r, err := redisPair.Get("myKey").Result()
	if err != nil {
		panic(err.Error())
	}

	log.Printf("result: %s\n", r)
}
