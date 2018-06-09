package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Init() {
	fmt.Println("~~~")
	http.HandleFunc("/", index)
	http.HandleFunc("/short", api_short)

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
			origin_url := string(result)
			// TODO: encode origin url here.
			fmt.Fprintln(resp, origin_url)
		}
	} else {
		resp.WriteHeader(404)
	}
}
