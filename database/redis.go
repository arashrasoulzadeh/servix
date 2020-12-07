package servix

import (
	"fmt"
	config "servix/config"

	"github.com/go-redis/redis"
)

var prefix = "servix_"

func Database() *redis.Client {
	redis_config := config.Parse().Redis
	return redis.NewClient(&redis.Options{
		Addr:     redis_config.Host + ":" + redis_config.Port,
		Password: redis_config.Password, // no password set
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
