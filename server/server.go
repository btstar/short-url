package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
)

func Init() {
	fmt.Print("[server] init redis...")
	redis.NewClient(&redis.Options{
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
				resp.WriteHeader(500)
			} else {
				// TODO: encode origin url here.
				fmt.Fprintln(resp, origin_url)
			}
		}
	} else {
		resp.WriteHeader(404)
	}
}
