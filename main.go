package main

import (
	"fmt"
	"net/http"
	"time"
)

var requestTotal int64

func main() {
	fmt.Println(time.Now(), "Server start")
	http.HandleFunc("/hello", sayHello)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	requestTotal++
	fmt.Println(time.Now(), "[", requestTotal, "]")
	go w.Write([]byte("hello"))
}
