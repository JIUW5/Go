package net

import (
	"net"
	"net/url"
)

/*
GetIp
解析输入的URL，并返回其对应的IP地址
*/
func GetIp(_url string) (string, error) {
	//2. 调用url.Parse方法解析_url，返回一个指向url.URL类型的指针pu和一个错误类型的值err
	pu, err := url.Parse(_url)
	if err != nil {
		return "", err
	}
	//4. 调用pu.Hostname方法获取主机名，赋值给host
	host := pu.Hostname()
	//5. 调用pu.Port方法获取端口号，赋值给port
	port := pu.Port()
	//6. 如果port为空字符串，设置port为"80"，如果pu.Scheme为"https"，设置port为"443"
	if port == "" {
		port = "80"
		if pu.Scheme == "https" {
			port = "443"
		}
	}
	//7. 调用net.ResolveTCPAddr方法解析TCP地址，返回一个指向net.TCPAddr类型的指针addr和一个错误类型的值err
	//	 net.JoinHostPort：将host和port拼接成host:port的形式
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
	//8. 如果err不为nil，返回空字符串和err
	if err != nil {
		return "", err
	}
	//9. addr.IP.String：获取IP地址并转换为字符串类型
	return addr.IP.String(), nil
}
