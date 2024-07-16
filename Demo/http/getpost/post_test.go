package getpost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

//详细解释一下这个代码
//1. 定义了一个结构体User，用来表示用户信息，包含了用户名和密码两个字段
//2. 定义了一个测试函数TestPostForm，用来测试POST请求，这个函数会向http://www.baidu.com发送一个POST请求，请求的参数是name=pibigstar&age=18
//3. 定义了一个测试函数TestHttppostJson，用来测试POST请求，这个函数会向http://www.baidu.com发送一个POST请求，请求的参数是一个JSON字符串，请求的Content-Type是application/json
//4. 定义了一个测试函数TestHttpDo，用来测试POST请求，这个函数会向http://www.baidu.com发送一个POST请求，请求的参数是一个JSON字符串，同时设置了Header和Cookie

// User 结构体,用来表示用户信息
type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func TestPostForm(t *testing.T) {
	//详细解释一下url.Values

	params := url.Values{}
	params.Add("name", "pibigstar")
	params.Add("age", "18")

	// Content-Type = application/x-www-form-urlencoded
	response, err := http.PostForm("http://www.baidu.com", params)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}

func TestHttppostJson(t *testing.T) {
	user := &User{
		UserName: "pibigstar",
		PassWord: "123456",
	}
	bs, _ := json.Marshal(user)
	data := bytes.NewReader(bs)

	resp, err := http.Post("http://www.baidu.com", "application/json", data)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}

// 有复杂请求，设置Header，cookie等需要使用这个
func TestHttpDo(t *testing.T) {
	client := &http.Client{}

	params := url.Values{}
	params.Add("name", "pibigstar")
	params.Add("age", "18")

	m := make(map[string]string)
	for key := range params {
		m[key] = params.Get(key)
	}
	bs, _ := json.Marshal(m)
	data := bytes.NewReader(bs)

	req, err := http.NewRequest("POST", "http://www.baidu.com", data)
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	req.Header.Set("Cookie", "name=anny")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))

}
