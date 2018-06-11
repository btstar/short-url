package main

import (
	"github.com/czsilence/short-url/server"
)

func main() {
	// set config
	cfg := server.DefaultServerConfig()
	cfg.Host = "0.0.0.0"
	cfg.Port = 9999
	cfg.IndexPath = "/"
	cfg.Url = "https://iurl.la"
	cfg.Redis = "127.0.0.1:16379"
	// init short-url server
	server.Init(cfg)
}
