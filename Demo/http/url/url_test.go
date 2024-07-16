package url

import (
	"net/url"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	v := url.Values{}
	v.Add("orgId", "123456")
	v.Add("userId", "Allenxu")
	// Encode函数用于将URL参数编码为字符串
	body := v.Encode()
	t.Log(body)
	// ParseQuery函数用于解析URL参数
	values, err := url.ParseQuery(body)
	if err != nil {
		t.Error(err)
	}
	t.Log(values.Get("orgId"))
}

func TestPathEscape(t *testing.T) {
	// PathEscape函数用于对URL路径进行编码，将URL路径中的特殊字符转换为%HH的形式
	encode := url.PathEscape("http://www.baidu.com?username=123&password=123")
	t.Log(encode)
	// PathUnescape函数用于对URL路径进行解码，将%HH的形式转换为特殊字符
	decode, err := url.PathUnescape(encode)
	if err != nil {
		t.Error(err)
	}
	t.Log(decode)
}
