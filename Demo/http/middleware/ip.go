package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

/*
	解释一下这些代码：
	1. 这个代码中定义了一个变量ipLimitMaps，它是一个map类型的变量，用来存储IP地址和对应的速率限制器
	2. 这个代码中定义了一个互斥锁mu，用来保护ipLimitMaps变量
	3. 这个代码中定义了两个变量rateLimit和rateMax，分别表示速率限制和速率限制的最大值
	4. GetIPLimiter函数用来获取IP地址对应的速率限制器
	5. IPRateLimit函数是一个中间件函数，用来对IP地址进行速率限制
	6. 这个中间件函数会检查IP地址对应的速率限制器，如果速率限制器允许请求通过，则调用传入的handler函数
	7. 如果速率限制器不允许请求通过，则返回429状态码
	这里的速率限制器是通过golang.org/x/time/rate包实现的，它可以限制单位时间内的请求次数
*/

var (
	ipLimitMaps = make(map[string]*rate.Limiter)
	mu          sync.Mutex
	rateLimit   = 1
	rateMax     = 5
)

func GetIPLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, ok := ipLimitMaps[ip]; ok {
		return limiter
	}

	limiter := rate.NewLimiter(rate.Limit(rateLimit), rateMax)
	ipLimitMaps[ip] = limiter

	return limiter

}

func IPRateLimit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := GetIPLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
