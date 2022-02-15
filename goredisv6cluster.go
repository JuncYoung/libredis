package libredis

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisGoV6Cluster(address, password string) *redis.ClusterClient {
	addrs := strings.Split(address, ",")
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       addrs,
		Password:    password,
		DialTimeout: time.Second * 10,
		//@TODO read write timeout
	})

	if str, err := rdb.Ping().Result(); err != nil {
		panic(fmt.Sprintf("address: %+v, p:%s,error:%s %s,\t %s:%s", address, password, err.Error(), str, address, password))
	}

	return rdb
}

//func (r *RedisGoV6Cluster) Close() {
//	if r == nil || r.client == nil {
//		return
//	}
//	r.client.Close()
//}
//
//func (r *RedisGoV6Cluster) Del(k string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	if err := r.client.Del(k).Err(); err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Expire(k string, timeout int) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	if err := r.client.Expire(k, time.Duration(timeout)*time.Second).Err(); err != nil {
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Get(k string) (int, string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, ""
//	}
//
//	cmdStr := r.client.Get(k)
//	ret, err := cmdStr.Result()
//	if err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK, ""
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, ""
//	}
//	return REDIS_OK, ret
//}
//
//func (r *RedisGoV6Cluster) Set(k string, v string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	if err := r.client.Set(k, v, REDIS_ERROR).Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Setex(k string, timeout int, v string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	if err := r.client.Set(k, v, time.Duration(timeout)*time.Second).Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//
//}
//
//func (r *RedisGoV6Cluster) Hdel(k string, k1 string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	if err := r.client.HDel(k, k1).Err(); err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//
//}
//
//func (r *RedisGoV6Cluster) HgetAll(k string) (int, map[string]string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, nil
//	}
//
//	cmdStr := r.client.HGetAll(k)
//	ret, err := cmdStr.Result()
//	if err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK, nil
//		}
//
//		return REDIS_ERROR, nil
//	}
//	return REDIS_OK, ret
//}
//
//func (r *RedisGoV6Cluster) Hget(k string, k1 string) (int, string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, ""
//	}
//
//	cmdStr := r.client.HGet(k, k1)
//	ret, err := cmdStr.Result()
//	if err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK, ""
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, ""
//	}
//	return REDIS_OK, ret
//}
//
//func (r *RedisGoV6Cluster) Hset(k string, field string, value string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	if err := r.client.HSet(k, field, value).Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Hsetnx(k string, field string, value string) (int, bool) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, false
//	}
//
//	isDone, err := r.client.HSetNX(k, field, value).Result()
//	if err != nil {
//
//		return REDIS_ERROR, false
//	}
//
//	return REDIS_OK, isDone
//}
//
//func (r *RedisGoV6Cluster) Setnx(k string, value string, expire time.Duration) (int, bool) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, false
//	}
//
//	isDone, err := r.client.SetNX(k, value, expire).Result()
//	if err != nil {
//
//		return REDIS_ERROR, false
//	}
//
//	return REDIS_OK, isDone
//}
//
//func (r *RedisGoV6Cluster) Lindex(k string, idx string) (int, string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, ""
//	}
//
//	i, err := strconv.Atoi(idx)
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, ""
//	}
//	cmdStr := r.client.LIndex(k, int64(i))
//	ret, err := cmdStr.Result()
//	if err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK, ""
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, ""
//	}
//	return REDIS_OK, ret
//}
//
//func (r *RedisGoV6Cluster) Llen(k string) (int, int) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, REDIS_OK
//	}
//
//	cmdStr := r.client.LLen(k)
//	ret, err := cmdStr.Result()
//	if err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK, REDIS_OK
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, REDIS_OK
//	}
//	return REDIS_OK, int(ret)
//}
//
//func (r *RedisGoV6Cluster) Lpop(k string) (int, string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, ""
//	}
//
//	cmdStr := r.client.LPop(k)
//	ret, err := cmdStr.Result()
//	if err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK, ""
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, ""
//	}
//	return REDIS_OK, ret
//}
//
//func (r *RedisGoV6Cluster) Lpush(k string, v string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	if err := r.client.LPush(k, v).Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Lrange(k string, start int, end int) (int, []string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, nil
//	}
//
//	strs, err := r.client.LRange(k, int64(start), int64(end)).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, nil
//	}
//	return REDIS_OK, strs
//}
//
//func (r *RedisGoV6Cluster) Rpop(k string) (int, string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, ""
//	}
//
//	cmdStr := r.client.RPop(k)
//	ret, err := cmdStr.Result()
//	if err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_OK, ""
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, ""
//	}
//	return REDIS_OK, ret
//
//}
//
//func (r *RedisGoV6Cluster) Rpush(k string, v string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	if err := r.client.RPush(k, v).Err(); err != nil {
//
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Sadd(k string, v string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	if err := r.client.SAdd(k, v).Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) SIsMember(k, element string) (bool, int) {
//	if r == nil || r.client == nil {
//		return false, REDIS_ERROR
//	}
//
//	boolCmd := r.client.SIsMember(k, element)
//	if err := boolCmd.Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return false, REDIS_ERROR
//	}
//	return boolCmd.Val(), REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Smembers(k string) (int, []string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, nil
//	}
//
//	strs, err := r.client.SMembers(k).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, nil
//	}
//	return REDIS_OK, strs
//}
//
//func (r *RedisGoV6Cluster) Srem(k string, v string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	if err := r.client.SRem(k, v).Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) SremBatch(k string, v []interface{}) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	if err := r.client.SRem(k, v...).Err(); err != nil {
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) SCard(k string) int64 {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	if count, err := r.client.SCard(k).Result(); err != nil {
//		if strings.Contains(err.Error(), "nil") {
//			return REDIS_ERROR
//		}
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return 0
//	} else {
//		return count
//	}
//}
//
//func (r *RedisGoV6Cluster) SPopN(k string, n int64) (int, []string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, nil
//	}
//
//	strs, err := r.client.SPopN(k, n).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, nil
//	}
//	return REDIS_OK, strs
//}
//
//// sorted set
//func (r *RedisGoV6Cluster) ZAddOne(k string, score float64, member string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	z := redis.Z{
//		Score:  score,
//		Member: member,
//	}
//	//v := make([]*redis.Z,0,1)
//	v := make([]redis.Z, 0, 1)
//	v = append(v, z)
//	if err := r.client.ZAdd(k, v...).Err(); err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) ZCard(k string) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//
//	count, err := r.client.ZCard(k).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, count
//	}
//
//	return REDIS_OK, count
//}
//
//func (r *RedisGoV6Cluster) ZRange(key string, start int64, stop int64) (int, []string) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, nil
//	}
//
//	strList, err := r.client.ZRange(key, start, stop).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, strList
//	}
//
//	return REDIS_OK, strList
//}
//
//func (r *RedisGoV6Cluster) ZRangeWithScores(k string, start int64, stop int64) (int, []redis.Z) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, nil
//	}
//
//	zList, err := r.client.ZRangeWithScores(k, start, stop).Result()
//	if err != nil {
//		return REDIS_ERROR, zList
//	}
//
//	return REDIS_OK, zList
//}
//
//func (r *RedisGoV6Cluster) ZRem(k string, members []interface{}) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//
//	count, err := r.client.ZRem(k, members...).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, count
//	}
//
//	return REDIS_OK, count
//}
//
//func (r *RedisGoV6Cluster) ZRank(k string, member string) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, -1
//	}
//
//	rank, err := r.client.ZRank(k, member).Result()
//	if err != nil {
//		if err.Error() != REDIS_NIL_STR {
//
//			return REDIS_ERROR, -1
//		}
//		return REDIS_OK, -1
//	}
//
//	return REDIS_OK, rank
//}
//
//func (r *RedisGoV6Cluster) ZRemrangebyscore(key string, min string, max string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	err := r.client.ZRemRangeByScore(key, min, max)
//	if err != nil {
//		return REDIS_ERROR
//	}
//
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Publish(channelKey string, data interface{}) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//
//	_, err := r.client.Publish(channelKey, data).Result()
//	if err != nil {
//
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Subscribe(channelKey []string) (*redis.PubSub, int) {
//	if r == nil || r.client == nil {
//		return nil, REDIS_ERROR
//	}
//	cmd := r.client.Subscribe(channelKey...)
//	return cmd, REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) BatchLpop(key string, start int64, end int64) ([]string, int) {
//	var result []string
//	if r == nil || r.client == nil {
//		return result, REDIS_ERROR
//	}
//	if start < 0 || start > end {
//		return result, REDIS_ERROR
//	}
//	pipe := r.client.Pipeline()
//	pipe.LRange(key, start, end-1)
//	pipe.LTrim(key, end, -1)
//	cmders, err := pipe.Exec()
//	if err != nil {
//		return result, REDIS_ERROR
//	}
//	for index, cmder := range cmders {
//		if index == 0 {
//			cmd := cmder.(*redis.StringSliceCmd)
//			result, err := cmd.Result()
//			if err != nil {
//				return result, REDIS_ERROR
//			}
//			return result, REDIS_OK
//		}
//	}
//
//	return result, REDIS_OK
//}
//
//func (r *RedisGoV6Cluster) Incr(key string) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.Incr(key).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) IncrBy(key string, value int64) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.IncrBy(key, value).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) IncrByFloat(key string, value float64) (int, float64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.IncrByFloat(key, value).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) Decr(key string) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.Decr(key).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) DecrBy(key string, value int64) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.DecrBy(key, value).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) HIncrBy(key, field string, value int64) (int, int64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.HIncrBy(key, field, value).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) HIncrByFloat(key, field string, value float64) (int, float64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.HIncrByFloat(key, field, value).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) ZIncr(key, Member string, score float64) (int, float64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	z := redis.Z{
//		Score:  score,
//		Member: Member,
//	}
//	res, err := r.client.ZIncr(key, z).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) ZIncrBy(key, Member string, score float64) (int, float64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	res, err := r.client.ZIncrBy(key, score, Member).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) ZIncrNX(key, Member string, score float64) (int, float64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	z := redis.Z{
//		Score:  score,
//		Member: Member,
//	}
//	res, err := r.client.ZIncrNX(key, z).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) ZIncrXX(key, Member string, score float64) (int, float64) {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR, 0
//	}
//	z := redis.Z{
//		Score:  score,
//		Member: Member,
//	}
//	res, err := r.client.ZIncrXX(key, z).Result()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR, 0
//	}
//	return REDIS_OK, res
//}
//func (r *RedisGoV6Cluster) Hmset(key string, fields map[string]interface{}) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	err := r.client.HMSet(key, fields).Err()
//	if err != nil {
//
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
//func (r *RedisGoV6Cluster) Lset(key string, index int64, value string) int {
//	if r == nil || r.client == nil {
//		return REDIS_ERROR
//	}
//	err := r.client.LSet(key, index, value).Err()
//	if err != nil {
//		time.Sleep(time.Millisecond * ERR_WAIT_INTERVAL_MS)
//		return REDIS_ERROR
//	}
//	return REDIS_OK
//}
