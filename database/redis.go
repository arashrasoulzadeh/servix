package servix

import (
	"fmt"

	"github.com/go-redis/redis"
)

var prefix = "servix_"

func Database() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func Set(key string, value string) {
	err := Database().Set(prefix+key, value, 0).Err()
	if err != nil {
		fmt.Errorf("cannot write to redis")
	}
}

func Incr(key string) {
	err := Database().Incr(prefix + key).Err()
	if err != nil {
		fmt.Errorf("cannot write to redis")
	}
}
func Get(key string, def string) string {
	val, err := Database().Get(prefix + key).Result()
	if err != nil {
		fmt.Errorf("cannot write to redis")
		return def
	}
	return val
}
