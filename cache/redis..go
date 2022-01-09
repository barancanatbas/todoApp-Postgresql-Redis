package cache

import (
	"encoding/json"
	"fmt"
	"todo/internal/config"

	"github.com/gomodule/redigo/redis"
)

func GetFromCache(key string, list interface{}) bool { // testler bittikten sonra buradaki fmt'ler kaldırılacak !
	// cache de var mı ?
	var client = config.Pool.Get()
	defer client.Close()
	value, err := redis.String(client.Do("get", key))
	if err != nil {
		//fmt.Println("->get hatası var")
		return false
	}
	err = json.Unmarshal([]byte(value), &list)
	if err != nil {
		fmt.Println("->json hatası var")
		return false
	}
	//fmt.Println("cache den aldık")

	return true
}

func SetFromCache(key string, list interface{}, ex uint64) {
	var client = config.Pool.Get()
	defer client.Close()
	jsondata, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json converter")
	}
	redis.Int64(client.Do("set", key, string(jsondata), "ex", ex))

}

func DeleteKeysFromPattern(pattern string) {
	var client = config.Pool.Get()
	defer client.Close()

	val, err := redis.Strings(client.Do("keys", pattern))
	if err == nil {
		for _, item := range val {
			client.Do("del", item)
		}
	}
}

func DelFromCache(key string) {
	var client = config.Pool.Get()
	client.Do("del", key)
	defer client.Close()
}
