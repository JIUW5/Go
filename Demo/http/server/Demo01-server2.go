package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
详细解释以下代码：
1. 代码中定义了一个变量count，用来记录访问的次数
2. 代码中定义了一个互斥锁mu，用来保护count变量
3. handler2函数会对count变量进行加1操作
4. counter函数会对count变量进行读操作
5. 由于count变量是一个共享变量，会被多个goroutine访问，所以需要使用互斥锁来保护这个变量
6. 互斥锁的Lock方法用来获取锁，Unlock方法用来释放锁
7. 互斥锁的Lock方法和Unlock方法之间的代码是临界区，同一时刻只能有一个goroutine进入临界区
8. 互斥锁的Lock方法和Unlock方法之间的代码是原子操作，不会被打断
9. 互斥锁的Lock方法和Unlock方法之间的代码是串行执行的，不会被并发执行
重点：
goroutine：goroutine是Go语言中的一个重要概念，它是一种轻量级的线程，由Go语言的运行时系统调度执行。
*/
var mu sync.Mutex
var count int

func main() {
	//这个函数的作用是启动一个web服务器，监听8000端口，当有请求到达 / 时，会调用handler2函数
	http.HandleFunc("/", handler2)
	// /counter 是一个相对路径，它是以 / 为根路径的，所以 /counter 也会触发handler2函数
	//当有请求到达 /counter 时，会调用counter函数
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 为什么，一次加3？
// 因为，浏览器会请求3次，分别是：/、/favicon.ico、/counter

func handler2(w http.ResponseWriter, r *http.Request) {

	//count变量是一个共享变量，会被多个goroutine访问，所以需要使用互斥锁来保护这个变量
	mu.Lock()
	count++
	//解锁
	mu.Unlock()
	fmt.Fprintf(w, "URL.PATH= %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	//保护count变量，防止在读取count变量的时候，有其他goroutine在修改count变量
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	//解锁
	mu.Unlock()
}
