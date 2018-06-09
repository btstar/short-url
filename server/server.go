package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
)

var redis_client *redis.Client

func Init() {
	fmt.Print("[server] init redis...")
	redis_client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	fmt.Print("OK\n")

	fmt.Print("[server] init http...")
	http.HandleFunc("/", index)
	http.HandleFunc("/short", api_short)

	fmt.Print("OK\n")
	fmt.Println("[server] start listening...")
	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		log.Fatal("http server start failed!", err)
	}
}

func index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "hello go!")
}

func api_short(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		log.Println("[server] handle api \"short\",", req.Method)
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			resp.WriteHeader(500)
		} else {
			origin_url := strings.TrimSpace(string(result))
			if len(origin_url) <= 0 {
				log.Fatal("[server] bad url data", origin_url)
				resp.WriteHeader(500)
			} else {
				ret := redis_client.Incr("short_url_id")
				if ret.Err() != nil {
					log.Fatal("[server] failed to prepare record", ret.Err())
					resp.WriteHeader(500)
				} else {
					url_id := ret.Val()
					ret := redis_client.HSet("url", string(url_id), origin_url)
					if ret.Err() != nil {
						log.Fatal("[server] failed to add record", ret.Err())
						resp.WriteHeader(500)
					} else {
						// TODO: 把id编码构造url下发
						fmt.Fprintln(resp, "id:", url_id)
					}
				}
			}
		}
	} else {
		resp.WriteHeader(404)
	}
}
