//go:generate go run -tags generate gen.go

package main

import (
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/schollz/jsonstore"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"

	redis2 "redis-client/redis"
)

type RedisPool struct {
	Name string
	pool *redis.Pool
}

type RedisPools struct {
	pools []RedisPool
}

type RedisProp struct {
	Name     string
	Host     string
	Port     int
	Password string
}

type RedisStoreData struct {
	KeyType string   `json:"keyType"`
	Ttl     string   `json:"ttl"`
	Value   []string `json:"value"`
}

var ks = new(jsonstore.JSONStore)
var configFile = "redis.json.gz"

func (r *RedisPools) loadRedisServer() []string {
	var err error
	ks, err = jsonstore.Open(configFile)
	if err != nil {
		ks = new(jsonstore.JSONStore)
	}
	keys := ks.Keys()
	for _, key := range keys {
		var redisProp RedisProp
		err = ks.Get(key, &redisProp)
		if err != nil {
			continue
		}
		r.initRedisClient(&redisProp)
	}
	return keys
}

func (r *RedisPools) initRedisClient(p *RedisProp) bool {
	log.Println("host:", p.Host)
	option := redis.DialPassword(p.Password)
	pool := &redis.Pool{
		MaxIdle:   10,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", strings.Join([]string{p.Host, ":", strconv.Itoa(p.Port)}, ""), option)
			if err != nil {
				panic(err)
			}
			return c, err
		},
	}
	r.pools = append(r.pools, *&RedisPool{
		p.Name,
		pool,
	})
	err := ks.Set(p.Name, p)
	if err != nil {
		panic(err)
	}
	if err = jsonstore.Save(ks, configFile); err != nil {
		panic(err)
	}
	return true
}

func selectPool(r *RedisPools, name string) *redis.Pool {
	for _, p := range r.pools {
		log.Println(p.Name)
		if p.Name == name {
			return p.pool
		}
	}
	return nil
}

func (r *RedisPools) selectDb(name string, db int) []string {
	pool := selectPool(r, name)
	if pool == nil {
		log.Println("pool init fail")
		return nil
	}
	c := pool.Get()
	// , "COUNT", "10000"
	arr, err1 := redis2.GetKeys(pool, db, "*")
	if err1 != nil {
		log.Println("scan failed:", err1)
		return nil
	}
	defer c.Close()
	return arr
}

func (r *RedisPools) valueByKey(name string, db int, key string) RedisStoreData {
	pool := selectPool(r, name)
	if pool == nil {
		log.Println("pool init fail")
		return *&RedisStoreData{Value: nil}
	}
	c := pool.Get()
	keyType := redis2.TypeKey(pool, key)
	log.Println("keyType:", keyType)
	switch keyType {
	case "hash":
		result, err := redis2.HScan(pool, key)
		if err != nil {
			return *&RedisStoreData{Value: nil}
		}
		return *&RedisStoreData{
			KeyType: "HASH",
			Value:   result,
		}
		break
	case "list":
		result, err := redis2.LRange(pool, key, 0, 999)
		if err != nil {
			return *&RedisStoreData{Value: nil}
		}
		return *&RedisStoreData{
			KeyType: "LIST",
			Value:   result,
		}
		break
	}
	arr, err1 := redis2.Get(pool, key)
	log.Println("get:", arr)
	if err1 != nil {
		log.Println("Get failed:", err1)
		return *&RedisStoreData{Value: nil}
	}
	defer c.Close()
	return *&RedisStoreData{Value: nil}
}

func main() {
	env := flag.String("env", "production", "")
	flag.Parse()
	log.Println(*env)

	var args []string
	args = append(args, "--disable-web-security")
	args = append(args, "--enable-file-cookies")
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 1024, 728, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	_ = ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	p := &RedisPools{}
	ui.Bind("loadRedisServer", p.loadRedisServer)
	ui.Bind("initRedisClient", p.initRedisClient)
	ui.Bind("selectDb", p.selectDb)
	ui.Bind("valueByKey", p.valueByKey)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	if *env == "dev" {
		ui.Load("http://localhost:3000")
	} else {
		ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	}

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	log.Println("exiting...")
}
