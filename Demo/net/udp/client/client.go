package main

import (
	"fmt"
	"net"
)

func main() {

	/*
		net.DialUDP函数用于建立一个UDP连接。
		它接收三个参数：网络类型，本地地址，和远程地址。
		这个函数返回一个net.UDPConn类型的连接对象和一个错误对象。
	*/
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8001,
	})

	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		/*
			使用conn.Write方法向UDP连接写入数据。
			写入的数据是一个字符串"hello,Allen"。
			这个方法返回写入的字节数和一个错误对象
		*/
		_, err := conn.Write([]byte("hello,Allen" + string(i)))
		if err != nil {
			fmt.Printf("failed to write msg,err:%v\n", err)
			break
		}

		var result [1024]byte
		/*
			程序使用conn.ReadFromUDP方法从UDP连接读取数据。
			这个方法返回读取的字节数，发送数据的远程地址，和一个错误对象。
		*/
		n, addr, err := conn.ReadFromUDP(result[:])
		if err != nil {
			fmt.Printf("failed to receiver msg,addr：%v ,err：%v\n", addr, err)
			break
		}
		fmt.Printf("reciver msg from：%v，msg：%s\n", addr, string(result[:n]))
	}
}
