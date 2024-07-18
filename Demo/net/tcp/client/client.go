package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//建立一个到本地8002端口的TCP连接。net.Dial函数返回一个net.Conn类型的连接对象和一个错误对象
	conn, err := net.Dial("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	//使用defer关键字来确保在main函数结束时，TCP连接会被关闭。
	defer conn.Close()

	//创建了一个新的读取器，用于从标准输入（通常是键盘）读取数据。
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please write your msg....")
	for {
		//从输入流中读取数据，直到遇到换行符 \n 为止。
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string from terminal, err: %v \n", err)
			break
		}
		//去掉读取到的数据两端的空格。
		data = strings.TrimSpace(data)
		//将读取到的数据写入到TCP连接中，通过[]byte()函数将字符串转换为字节切片。
		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Printf("failed to write server: %v \n", err)
		}
		if data == "exit" {
			break
		}

		var result [1024]byte
		n, err := conn.Read(result[:])
		if err != nil {
			fmt.Printf("failed to read result from server, err: %v \n", err)
			continue
		}
		fmt.Printf("read result from server: %s \n", string(result[:n]))
	}
}
