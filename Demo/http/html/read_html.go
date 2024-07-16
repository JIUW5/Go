package main

import "net/http"

func main() {
	//创建了一个http.Dir对象，它表示一个文件系统子树。它指向了项目中的Demo/http/html/root目录。
	dir := http.Dir("Demo/http/html/root")
	/*
		创建了一个新的http.FileServer，它使用http.Dir对象（即dir）作为文件系统接口。
		http.FileServer是一个HTTP处理器，它将HTTP请求转换为对dir的文件系统操作。
	*/
	staticHandler := http.FileServer(dir)

	/*
		将http.FileServer处理器注册到HTTP服务器的根路径（"/"）。
		StripPrefix函数将请求的URL路径中的前缀（在这个例子中也是"/"）去掉，然后将请求传递给staticHandler处理。
		例如，一个对/index.html的请求就会被转换为对index.html文件的请求。
	*/
	http.Handle("/", http.StripPrefix("/", staticHandler))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
