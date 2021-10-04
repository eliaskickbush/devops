package cache

import (
	"github.com/go-redis/redis"
	"log"
)

type Dao struct {
	RedisClient *redis.Client
	Logger *log.Logger
}

func (d Dao) SetValue(key string, value interface{}) error {
	status := d.RedisClient.Set(key,value, 0)
	if status.Err() != nil {
		d.Logger.Printf("there was an error setting key value pair (%s,%v): %s", key,value,status.Err().Error())
		return ErrSettingValue
	}
	return nil
}

func (d Dao) GetValue(key string) (string, error) {
	stringCommand := d.RedisClient.Get(key)
	if stringCommand.Err() != nil {
		d.Logger.Printf("there was an error getting value for key %s: %s\n", key, stringCommand.Err().Error())
		return "", ErrGettingValue
	}
	return stringCommand.Val(), nil
}

func (d Dao) ClearValue(key string) error {
	intCommand := d.RedisClient.Del(key)
	if intCommand.Err() != nil {
		d.Logger.Printf("there was an error deleting value for key %s: %s\n", key, intCommand.Err().Error())
		return ErrDeletingValue
	}
	return nil
}



