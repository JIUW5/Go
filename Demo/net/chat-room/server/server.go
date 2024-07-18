package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// 存放客户端连接切片
var clients []net.Conn

var (
	host   = "localhost"
	port   = "8888"
	remote = fmt.Sprintf("%s:%s", host, port)
	data   = make([]byte, 1024)
)

func main() {

	fmt.Print("Initiating server........")
	/*
		在指定的地址（由host和port变量组成）上启动一个TCP服务器。
		net.Listen函数返回一个net.Listener类型的监听器对象和一个错误对象。
	*/
	listener, err := net.Listen("tcp", remote)
	defer listener.Close()
	if err != nil {
		log.Printf("Error when listen:%s \n", err.Error())
		/*
			os.Exit(-1) 是 Go 语言中用于立即终止程序运行的函数调用。os.Exit 函数接受一个整数参数，这个参数被称为退出状态码。
			通常情况下，当程序正常退出时，会返回 0 作为退出状态码；当程序非正常退出时，会返回非 0 的退出状态码。
		*/
		os.Exit(-1)
	}
	fmt.Println("Ok!")

	for {
		/*
			等待并接受来自客户端的连接。listener.Accept函数返回一个net.Conn类型的连接对象和一个错误对象。
		*/
		conn, err := listener.Accept()

		defer conn.Close()
		if err != nil {
			log.Printf("Error when get connect:%s \n", err.Error())
		}
		//将新的客户端连接添加到clients切片中
		clients = append(clients, conn)

		// 为每一个连接分配一个goroutine
		//这是一个匿名函数，相关格式可以上网查阅
		/*
			(conn) 是 Go 语言中的函数参数。
			在这个例子中，它是传递给匿名函数的参数，表示一个客户端的连接。
			这个连接是 net.Conn 类型，它是 Go 标准库中的一个接口类型，用于表示网络连接。
			go func(con net.Conn) {...} 定义了一个匿名函数，这个函数接受一个 net.Conn 类型的参数 con。
			然后，(conn) 是对这个匿名函数的调用，传入的实参是 `conn`，它是从 listener.Accept() 返回的客户端连接。
		*/
		go func(con net.Conn) {

			//RemoteAddr() 远程地址
			fmt.Printf("New Connection : %s \n", con.RemoteAddr())
			// 得到客户端的名字
			length, err := con.Read(data)
			if err != nil {
				log.Printf("client : %s quit", con.RemoteAddr())
				con.Close()
				return
			}
			name := string(data[:length])
			connName := name + " enter the room"
			// 通知其他客户端 我进来了
			notify(con, connName)

			// 监听其他客户端的消息
			for {
				length, err := con.Read(data)
				if err != nil {
					log.Printf("client : %s quit", con.RemoteAddr())
					con.Close()
					disconnect(con, name)
					return
				}
				res := string(data[:length])
				//表示该客户端说了什么
				msg := fmt.Sprintf("%s said: %s", name, res)
				fmt.Println(msg)
				//表示服务器对客户端的回应
				res = fmt.Sprintf("You said:%s", res)
				con.Write([]byte(res))
				notify(con, msg)
			}
		}(conn)
	}
}

// 群发其他客户端
func notify(conn net.Conn, msg string) {
	for _, con := range clients {
		if con.RemoteAddr() != conn.RemoteAddr() {
			con.Write([]byte(msg))
		}
	}
}

// 断开连接，并通知其他客户端
func disconnect(conn net.Conn, name string) {
	for index, con := range clients {
		if con.RemoteAddr() == conn.RemoteAddr() {
			// 移除此 客户端_
			clients = append(clients[:index], clients[index+1:]...)
			disMsg := fmt.Sprintf("%s %s", name, "left the room")
			notify(con, disMsg)
		}
	}
}
