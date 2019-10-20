package router

import (
	"fmt"
	"net/http"
)

func RequestCountingMiddleware(next http.Handler) http.Handler {
	mutex.Lock()
	pendingRequests++
	mutex.Unlock()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		mutex.Lock()
		pendingRequests--
		mutex.Unlock()
	})
}

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Got some request on", r.URL.String())
		next.ServeHTTP(w, r)
	})
}
