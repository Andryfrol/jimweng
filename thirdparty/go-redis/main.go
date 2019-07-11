package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func NewRedisClient() redisDAO {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rO := &redisObject{
		ro: client,
	}
	return rO
}

func main() {
	redisClient := NewRedisClient()
	redisClient.ping()
	// fmt.Printf("%v\n", pong)
	redisClient.setValue("a", "jim")
	redisClient.queryValue("a")
}

type redisObject struct {
	ro *redis.Client
}

type redisDAO interface {
	ping()                   // expect to print out pong on the console
	setValue(string, string) // expect to set value into redis
	queryValue(string)
	deleteValue(string)
}

func (r *redisObject) ping() {
	pong, _ := r.ro.Ping().Result()
	fmt.Println(pong)
}

func (r *redisObject) setValue(key string, value string) {
	err := r.ro.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r *redisObject) queryValue(key string) {
	if val, err := r.ro.Get(key).Result(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Query key '%s', get return value '%s'\n", key, val)
	}
}

func (r *redisObject) deleteValue(key string) {

}

// val2, err := client.Get("key2").Result()
// if err == redis.Nil {
// 	fmt.Println("key2 does not exist")
// } else if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println("key2", val2)
// }
// Output: key value
// key2 does not exist
