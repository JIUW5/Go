package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	/*
		net.Listen函数用于在指定的地址上监听TCP连接。
		它接收两个参数：网络类型和本地地址。
		这个函数返回一个net.Listener类型的监听对象和一个错误对象
	*/
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	fmt.Println("tcp server is running on :8002.......")

	for {

		conn, err := listener.Accept()
		if err != nil {
			//如果在监听时发生错误，程序会调用panic函数抛出一个运行时错误，并终止执行
			panic(err)
		}
		//启动一个新的goroutine（Go语言中的轻量级线程）来处理一个新的TCP连接。
		go process(conn)
	}
}

func process(conn net.Conn) {
	/*
		确保无论process函数如何结束（正常结束或因为错误而结束），conn.Close()方法都会被执行，从而确保TCP连接被正确关闭
	*/
	defer conn.Close()
	for {
		var data [1024]byte
		/*
			从TCP连接中读取数据。接收一个字节切片作为参数，并尝试从连接中读取数据填充这个字节切片。
		*/
		n, err := conn.Read(data[:])
		if err != nil && err != io.EOF {
			fmt.Printf("failed to read data from client, err: %v \n", err)
			break
		}
		str := string(data[:n])
		if str == "exit" {
			fmt.Println("client exit....")
			break
		}

		fmt.Printf("read data from client: %s \n", str)

		// 回复ACK
		conn.Write([]byte(fmt.Sprintf("%s OK!", str)))
	}
}
