package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	// Port service.port
	Port int
	// RedisAddr redis.connect.addr
	RedisAddr string
	// RedisPwd redis.connect.password
	RedisPwd string
	// RedisDB redis.connect.db
	RedisDB int
	// RedisClient redis.client
	RedisClient *redis.Client
)

func init() {
	flag.IntVar(&Port, "port", 3001, "server.port")
	flag.StringVar(&RedisAddr, "addr", "localhost:6379", "redis.addr")
	flag.StringVar(&RedisPwd, "pwd", "", "redis.password")
	flag.IntVar(&RedisDB, "db", 0, "redis.db")
}

func main() {

	flag.Parse()
	redisConn()

	addr := fmt.Sprintf(":%d", Port)
	fmt.Printf("Server start at %s\n", addr)
	r := mux.NewRouter()
	r.HandleFunc("/{k}", Handler).Methods("GET", "POST", "DELETE")
	http.ListenAndServe(addr, r)
}

// Handler 回调
func Handler(w http.ResponseWriter, r *http.Request) {

	// init
	mtd := r.Method
	vars := mux.Vars(r)
	k, ok := vars["k"]
	if !ok {
		fmt.Fprintln(w, "fail to request: without key\n")
	}

	fmt.Printf("Handler[%s](%s)\n", mtd, k)

	switch mtd {
	case "GET":
		v, err := redisGet(k)
		if nil != err {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
		}
		fmt.Fprintln(w, v)
	case "DELETE":
		err := redisDel(k)
		if nil != err {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
		}
	case "POST":
		err := redisSet(k, getParam(r, "v"))
		if nil != err {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
		}
	default:
		fmt.Fprintf(w, "fail to request: not supported type[%s]\n", mtd)
	}

}

func getParam(r *http.Request, k string) string {

	args := r.URL.Query()
	v := ""

	// check
	vs, ok := args[k]
	if ok && len(vs) > 0 {
		v = vs[0]
	}

	return v
}

func redisConn() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPwd,
		DB:       RedisDB,
	})

	pong, err := RedisClient.Ping().Result()
	if nil != err {
		panic("fail to connect to redis[" + RedisAddr + "]: " + err.Error())
	}
	fmt.Printf("redis.ping= %s\n", pong)
}

func redisSet(k, v string) error {

	err := RedisClient.Set(k, v, -1).Err()
	if nil != err {
		fmt.Errorf("fail to set key[%s] to %s: %s\n", k, v, err.Error())
		return err
	}
	fmt.Printf("Set[%s]= %s\n", k, v)
	return nil
}

func redisDel(k string) error {

	err := RedisClient.Del(k).Err()
	if nil != err {
		fmt.Errorf("fail to del key[%s]: %s\n", k, err.Error())
		return err
	}
	fmt.Printf("Del[%s]\n", k)
	return nil
}

func redisGet(k string) (string, error) {

	v, err := RedisClient.Get(k).Result()
	if nil != err {
		fmt.Errorf("fail to get key[%s]: %s\n", err.Error())
		return "", err
	}

	fmt.Printf("Get[%s]= %s\n", k, v)
	return v, nil
}
