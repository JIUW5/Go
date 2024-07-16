package middleware

import "net/http"

/*
Auth 验证token
1. 这个函数接受一个http.HandlerFunc类型的参数，返回一个http.HandlerFunc类型的函数
2. 这个函数会检查请求中的token参数，如果token的值是"pi"，则调用传入的handler函数
3. 如果token的值不是"pi"，则返回403状态码
*/
func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("token")
		if name == "pi" {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
}
