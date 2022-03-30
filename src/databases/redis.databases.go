package databases

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/fecamp-cu/fe-2021-backend-settings/src/configs"
	"github.com/go-redis/redis/v8"
)

var r RedisClient
var lockRedis sync.Once
var ctx = context.Background()

func initRedis() {
	rdb := redis.NewClient(getRedisOptions())
	r = RedisClient{rdb}
}

type RedisClient struct {
	client *redis.Client
}

func GetRedis() RedisClient {
	lockRedis.Do(initRedis)
	return r
}

func getRedisOptions() *redis.Options {
	redisConfigs := configs.GetConfigs().Redis
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfigs.Host, redisConfigs.Port),
		Password: redisConfigs.Password,
		DB:       0,
	}
}

func (rc *RedisClient) Get(key string, data interface{}) error {
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), data)
}

func (rc *RedisClient) Set(key string, data interface{}) error {
	p, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return rc.client.Set(ctx, key, p, 0).Err()
}

func (rc *RedisClient) Delete(key string) error {
	return rc.client.Del(ctx, key).Err()
}
