package main

import (
	"fmt"
	"log"
	"net/http"
)

// 这个代码是一个简单的web服务器，它会打印出请求的URL路径
func main() {

	//这个函数的作用是启动一个web服务器，监听8000端口，当有请求到达时，会调用handler函数
	//HandleFunc函数的第一个参数是URL路径，第二个参数是一个函数，这个函数会处理所有到达这个URL的HTTP请求
	http.HandleFunc("/", handler1)
	//log.Fatal函数会打印错误信息并退出程序
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// w 是一个用来写入HTTP响应的接口，作用：将字符串写入到HTTP响应中，这样客户端就能看到这个响应了
// r 是一个指向http.Request结构的指针，这个结构中包含了这个HTTP请求的所有的信息，比如请求的URL是什么
func handler1(w http.ResponseWriter, r *http.Request) {
	//Fprintf函数可以将格式化的字符串写入到io.Writer接口中，这里我们使用fmt.Fprintf将字符串写入到w中
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
