package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 1 实现一个 http server
// 2 实现一个 handler：hello
// 3 实现中间件的功能：1）记录请求的 URL 和方法 2）记录请求的网络地址 3）记录方法的执行时间

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("tracing start 记录请求的 URL 和方法: %s, %s", r.URL, r.Method)
		next.ServeHTTP(w, r)
		log.Println("tracing end")
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("logging strat 记录请求的网络地址: %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Println("logging end")
	})
}

func timeRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("timeRecording start")
		startTime := time.Now()
		next.ServeHTTP(w, r)
		elapsedTime := time.Since(startTime)
		log.Printf("记录方法的执行时间: %s", elapsedTime)
		log.Println("timeRecording end")
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello")
}

func main() {
	http.Handle("/", tracing(logging(timeRecording(http.HandlerFunc(hello)))))
	_ = http.ListenAndServe(":8080", nil)
}
