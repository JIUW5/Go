package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//解释这段代码：
//1. 定义了一个Body结构体，这个结构体有两个字段，Name和Age

type Body struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 2. 定义了一个测试函数TestHttpTest
func TestHttpTest(t *testing.T) {
	//3. 在测试函数中，创建了一个http.ServeMux对象，这个对象是一个路由器，用于匹配请求的URL路径，并调用相应的处理函数
	//&:取地址符号，表示mux是一个指针类型
	mux := &http.ServeMux{}

	//4. 通过mux.HandleFunc函数，注册了两个处理函数，分别是/get和/post
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		t.Log(r.FormValue("name"))
		//5. 在/get处理函数中，通过r.FormValue函数，获取请求的name参数，并将name参数写入到http.ResponseWriter中
		w.Write([]byte(r.FormValue("name")))
	})

	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		var body Body
		if r.Body != nil {
			//6. 在/post处理函数中，通过json.NewDecoder(r.Body).Decode(&body)函数，将http.Request.Body字段的内容解析为JSON格式，并将解析的结果存储到body变量中
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
		}
		t.Logf("%+v", body)
		body.Age = 20
		//7. 在/post处理函数中，通过json.Marshal(body)函数，将body变量转换为JSON格式的字节切片，并将字节切片写入到http.ResponseWriter中
		bs, _ := json.Marshal(body)
		w.Write(bs)
	})
	//8. 通过httptest.NewServer函数，创建了一个测试服务器
	ts := httptest.NewServer(mux)
	//9. 通过http.Get函数，向测试服务器发送一个GET请求，请求的URL是/ts.URL + "/get?name=allenxu"
	response, err := http.Get(ts.URL + "/get?name=allenxu")
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	t.Log(string(bs))

	body := &Body{Name: "Hello"}
	bs, _ = json.Marshal(body)
	//10. 通过http.Post函数，向测试服务器发送一个POST请求，请求的URL是ts.URL+"/post"，请求的Body是一个JSON格式的字节切片
	response, err = http.Post(ts.URL+"/post", "application/json", bytes.NewReader(bs))
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	//11. 通过ioutil.ReadAll函数，读取响应的Body字段，并将读取的结果打印到控制台中
	bs, _ = ioutil.ReadAll(response.Body)
	t.Log(string(bs))
}
