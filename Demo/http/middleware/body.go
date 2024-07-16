package middleware

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
	BodyLimit

1. 这个代码定义了一个BodyLimit函数，它接受一个http.HandlerFunc类型的参数，返回一个http.HandlerFunc类型的函数
2. 这个函数会检查请求的方法是否是POST，如果不是POST方法，则返回405状态码
3. 这个函数会读取请求的body数据，如果body数据的长度超过128字节，则返回400状态码
4. 如果body数据的长度没有超过128字节，则调用传入的handler函数
*/
func BodyLimit(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var maxLength int64 = 128
		var body = make(map[string]interface{})

		err := json.NewDecoder(io.LimitReader(r.Body, maxLength)).Decode(&body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			if err == io.EOF {
				w.Write([]byte("No body data"))
			} else {
				w.Write([]byte("Body length illegal"))
			}
			return
		}
		handler.ServeHTTP(w, r)
	}
}
