package server

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	base62 "github.com/pilu/go-base62"
)

var redis_client *redis.Client
var config ServerConfig

func Init(cfg *ServerConfig) {

	if cfg == nil {
		cfg = DefaultServerConfig()
	}
	config = *cfg

	log.Print("[server] init redis...")
	redis_client = redis.NewClient(&redis.Options{
		Addr:     config.Redis,
		Password: "",
		DB:       0,
	})
	log.Print("OK\n")

	log.Print("[server] init http...")
	http.HandleFunc(config.IndexPath, index)

	log.Print("OK\n")
	log.Println("[server] start listening...")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), nil)
	if err != nil {
		log.Fatal("http server start failed!", err)
	}
}

func index(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		// 生成短链接
		api_short(resp, req)
		return
	}

	if req.Method != "GET" {
		resp.WriteHeader(404)
		return
	}

	path := strings.Replace(req.URL.Path, config.IndexPath, "", 1)
	log.Println("req:", path)
	if path == "/" {
		resp.WriteHeader(404)
		return
	}

	// path = path[1:]
	ret := redis_client.HGet("url", strconv.Itoa(base62.Decode(path)-12345))
	if ret.Err() != nil {
		log.Println("[server] no url matched:", path)
		resp.WriteHeader(404)
	} else {
		origin_url := ret.Val()
		log.Printf("[server] url, %s => %s", path, origin_url)
		http.Redirect(resp, req, origin_url, 301)
	}
}

func api_short(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		log.Println("[server] handle api \"short\",", req.Method)
		result, err := ioutil.ReadAll(req.Body)
		if err != nil {
			resp.WriteHeader(500)
		} else if origin_url := strings.TrimSpace(string(result)); len(origin_url) <= 0 {
			log.Println("[server] bad url data", origin_url)
			resp.WriteHeader(500)
		} else {
			hash := getUrlHash(origin_url)

			if url, ex := checkExist(hash); ex {
				writeShortId(resp, url)
			} else {
				ret := redis_client.Incr("short_url_id")
				if ret.Err() != nil {
					log.Println("[server] failed to prepare record", ret.Err())
					resp.WriteHeader(500)
				} else {
					url_id := ret.Val()
					ret := redis_client.HSet("url", string(url_id), origin_url)
					if ret.Err() != nil {
						log.Println("[server] failed to add record", ret.Err())
						resp.WriteHeader(500)
					} else {
						// FIXME: 潜在的溢出，暂时不用管，64位系统上应该没有关系
						url := base62.Encode(int(url_id + 12345))
						redis_client.HSet("origin_url", hash, url)
						writeShortId(resp, url)
					}
				}
			}
		}
	} else {
		resp.WriteHeader(404)
	}
}

func checkExist(hash string) (url string, ex bool) {
	if ret := redis_client.HGet("origin_url", hash); ret.Err() != nil {
		ex = false
	} else if url = ret.Val(); len(url) <= 0 {
		ex = false
	} else {
		ex = true
	}
	return
}

func getUrlHash(origin_url string) string {
	ctx := md5.New()
	ctx.Write([]byte(origin_url))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 把id编码构造url下发
func writeShortId(resp http.ResponseWriter, url string) {
	fmt.Fprint(resp, makeUrl(&url))
}

// 构造url
func makeUrl(id *string) string {
	return fmt.Sprintf("%s/%s", config.Url, *id)
}
