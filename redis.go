package commonDB

// redis is a database component for key/value storing

import (
	"encoding/json"

	"gopkg.in/redis.v3"
)

// RedisClient is our redis DB that implements the Component Interface
type RedisClient struct {
	*redis.Client
}

// NewRedisClient Creates client and opens a connection to the database
func NewRedisClient() *RedisClient {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rClient.Ping().Result()
	if err != nil {
		panic("Please check Redis connection. Error: " + err.Error())
	}
	r := &RedisClient{
		rClient,
	}
	return r
}

// StoreStruct takes any struct, serializes it into json and stores in redis
func (r *RedisClient) StoreStruct(key string, object interface{}) error {
	// serialized object
	encObj, err := json.Marshal(object)
	if err != nil {
		return err
	}
	_, err = r.Set(key, string(encObj), 0).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetStruct takes a key and tries to construct an object.
func (r *RedisClient) GetStruct(key string, object interface{}) error {
	strObj, err := r.Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(strObj), object)
}
