package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
这段代码，是干嘛用的？详细跟我介绍一下！
1. 通过http.HandleFunc函数，将handler3函数与/路径绑定
2. 通过http.ListenAndServe函数，监听8000端口
3. handler3函数，接收两个参数，一个是http.ResponseWriter，一个是http.Request
4. handler3函数，通过fmt.Fprintf函数，将请求的方法、URL、协议版本，写入到http.ResponseWriter中
5. handler3函数，通过for循环，遍历请求的Header，将请求的Header写入到http.ResponseWriter中
6. handler3函数，通过fmt.Fprintf函数，将请求的Host、RemoteAddr，写入到http.ResponseWriter中
7. handler3函数，通过r.ParseForm函数，解析请求的Form字段
8. handler3函数，通过for循环，遍历请求的Form字段，将请求的Form字段写入到http.ResponseWriter中
*/
func main() {
	http.HandleFunc("/", handler3)
	//nil: 表示使用标准库中的DefaultServeMux,什么意思？
	//DefaultServeMux: 是一个ServeMux类型的变量，是一个路由器
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	//w: 用于写入响应的http.ResponseWriter
	fmt.Fprintf(w, "请求方法： %s,请求路径：%s,http协议：%s\n", r.Method, r.URL, r.Proto)
	//k: 请求头的key,v: 请求头的value
	//解释一下
	//1. r.Header是一个map[string][]string类型
	//2. range是一个迭代器，用于遍历r.Header
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	//Host: 请求的主机
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	//RemoteAddr: 请求的远程地址
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	//1. err := r.ParseForm()，解析请求的Form字段
	//2. err != nil，说明解析失败
	//3. log.Print(err)，打印错误信息
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
