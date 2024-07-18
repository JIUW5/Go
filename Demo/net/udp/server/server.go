package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	/*
		net.ListenUDP函数用于在指定的地址上监听UDP连接。
		它接收两个参数：网络类型和本地地址（在这个例子中是一个指向net.UDPAddr类型的指针，IP地址为0.0.0.0，端口号为8001）。
		这个函数返回一个net.UDPConn类型的监听对象和一个错误对象。
	*/
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8001,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("udp server listening......")

	for {
		var data [1024]byte
		/*
			使用listen.ReadFromUDP方法从UDP连接读取数据。
			这个方法返回读取的字节数，发送数据的远程地址，和一个错误对象
		*/
		n, addr, err := listen.ReadFromUDP(data[:])
		/*
			如果在读取数据时发生错误，并且错误不是io.EOF（表示没有更多的数据可以读取），程序会打印错误信息，并跳出循环
		*/
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		/*
			程序启动一个新的goroutine（轻量级线程）
			在这个goroutine中，程序会打印出发送数据的远程地址和接收到的数据，然后向远程地址发送一个响应消息
		*/
		go func() {
			str := string(data[:n])
			fmt.Printf("read data: %s, addr：%v, count：%d \n", str, addr, n)

			n, err = listen.WriteToUDP([]byte(fmt.Sprintf("%s OK!", str)), addr)
			if err != nil {
				fmt.Printf("failed to write udp,addr:%v\n", addr)
			}
		}()
	}
}
