package main

import (
	"Basic/Demo/http/middleware"
	"fmt"
	"net/http"
)

func main() {
	//创建一个ServeMux类型的变量，这个变量是一个路由器
	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.BodyLimit(hello))
	mux.HandleFunc("/admin", middleware.Auth(hello))

	//go 是什么？go是一个关键字，用于启动一个goroutine
	go fmt.Println("server starting...")
	if err := http.ListenAndServe(":8081", middleware.IPRateLimit(mux)); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")

}
