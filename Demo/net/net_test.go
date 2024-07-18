package net_test

import (
	"Basic/Demo/net"
	"testing"
)

/*
TestGetIp
1. TestGetIp函数接收一个指向testing.T类型的指针t
2. 调用net.GetIp方法，传入"http://www.baidu.com"，返回一个字符串类型的值和一个错误类型的值
*/
func TestGetIp(t *testing.T) {
	ip, err := net.GetIp("http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}

	t.Log(ip)
}
