package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	os.Setenv("VERSION", "1.0")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// define a server
	srv := &http.Server{Addr: ":80"}
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)

	// run a server
	go func(ctx context.Context) {
		log.Println("server is running...")
		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("server run failed:%+v", err)
		}
		log.Println("server is stoppping...")
	}(ctx)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	<-done // stop here until get signal above
	cancel()

	// stop server
	err := srv.Shutdown(context.TODO())
	if err != nil {
		log.Printf("server shutdown failed:%+v", err)
	}
	log.Println("server is stopped.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}

	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		w.Header().Add(k, strings.Join(v, ",")) // F12 看不到
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	fmt.Fprintf(w, "VERSION = %q\n", os.Getenv("VERSION"))
	w.Header().Add("VERSION", os.Getenv("VERSION"))

	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Printf("IP = %q, Status = %q\n", r.Host, "200")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 4.当访问 localhost/healthz 时，应返回200
	fmt.Fprintf(w, "200 OK")
}
