package redis

import (
	"errors"
	"go-project/config"
	"go-project/utils/log"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	cache    *CacheRedis
	ErrRedis = errors.New("redis is unavalaible")
)

type CacheRedis struct {
	Client  *redis.Client
	isReady bool
}

func InitCacheConnection() error {
	err := InitConnectionCache()
	if err != nil {
		return err
	}
	return nil
}

func InitConnectionCache() error {
	cache = new(CacheRedis)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.AppEnv.RedisAddress,
		Password: config.AppEnv.RedisPassword,
		DB:       config.AppEnv.RedisDb,
	})

	ping, err := redisClient.Ping().Result()
	if err != nil {
		cache.isReady = false
		return err
	}

	log.Log.Info("redis ping : ", ping)

	cache.Client = redisClient
	cache.isReady = true

	return err
}

func GetCacheConnection() (*CacheRedis, error) {
	return cache.getConnection()
}

func (cRedis *CacheRedis) getConnection() (*CacheRedis, error) {
	if !cRedis.isReady {
		return nil, ErrRedis
	}
	_, err := cRedis.Client.Ping().Result()
	if err != nil {
		cRedis.isReady = false
		return nil, err
	}
	return cRedis, nil
}

func CloseCacheConnection() {
	logrus.Println("Closing Cache connection...")
	cache.Client.Close()
}

func (cRedis *CacheRedis) GetValue(key string) ([]byte, error) {
	return cRedis.Client.Get(key).Bytes()
}

func (cRedis *CacheRedis) GetValueString(key string) (string, error) {
	return cRedis.Client.Get(key).Result()
}

func (cRedis *CacheRedis) GetValueInt(key string) (int, error) {
	return cRedis.Client.Get(key).Int()
}
