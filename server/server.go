package server

import (
	"fmt"
	"log"
	"net/http"
)

func Init() {
	fmt.Println("~~~")
	http.HandleFunc("/", index)

	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		log.Fatal("http server start failed!", err)
	}
}

func index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "hello go!")
}
