package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

func Ping(p *redis.Pool) error {

	conn := p.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func TypeKey(p *redis.Pool, key string) string {
	conn := p.Get()
	defer conn.Close()
	data, err := redis.String(conn.Do("TYPE", key))
	if err != nil {
		return ""
	}
	return data
}

func HScan(p *redis.Pool, key string) ([]string, error) {
	conn := p.Get()
	defer conn.Close()
	data, err := redis.Values(conn.Do("HSCAN", key, "0", "COUNT", "10000"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return parseScanResults(data)
}

func parseScanResults(results []interface{}) ([]string, error) {
	if len(results) != 2 {
		return []string{}, nil
	}

	_, err := strconv.ParseInt(string(results[0].([]byte)), 10, 64)
	if err != nil {
		return nil, err
	}

	keyInterfaces := results[1].([]interface{})
	keys := make([]string, len(keyInterfaces))
	for index, keyInterface := range keyInterfaces {
		keys[index] = string(keyInterface.([]byte))
	}
	return keys, nil
}

func Get(p *redis.Pool, key string) ([]byte, error) {
	conn := p.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func Set(p *redis.Pool, key string, value []byte) error {

	conn := p.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func Exists(p *redis.Pool, key string) (bool, error) {

	conn := p.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func Delete(p *redis.Pool, key string) error {
	conn := p.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func GetKeys(p *redis.Pool, num int, pattern string) ([]string, error) {
	conn := p.Get()
	defer conn.Close()
	conn.Do("SELECT", num)
	iter := 0
	var keys []string
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

func Incr(p *redis.Pool, counterKey string) (int, error) {

	conn := p.Get()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", counterKey))
}
