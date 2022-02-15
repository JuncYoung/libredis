package libredis

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type ConfigOption func(fo *redis.FailoverOptions)

func NewRedisGoV6Sentinel(addrs, password, masterName string, opts ...ConfigOption) *redis.Client {
	addrlist := strings.Split(addrs, ",")
	failoverOptions := &redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: addrlist,
		Password:      password,
		DialTimeout:   time.Second * 30,
		ReadTimeout:   time.Second * 30,
		WriteTimeout:  time.Second * 30,
		MaxRetries:    3,
	}
	for _, opt := range opts {
		opt(failoverOptions)
	}

	rdb := redis.NewFailoverClient(failoverOptions)

	if str, err := rdb.Ping().Result(); err != nil {
		panic(fmt.Sprintf("address: %+v, p:%s,error:%s %s,\t %s:%s", addrs, password, err.Error(), str, addrs, password))
	}
	return rdb
}

//MaxRetries implements ...
func MaxRetries(count int) ConfigOption {
	return func(fo *redis.FailoverOptions) {
		fo.MaxRetries = count
	}
}

//DBIndex implements ...
func DBIndex(dbIndex int) ConfigOption {
	return func(fo *redis.FailoverOptions) {
		fo.DB = dbIndex
	}
}
