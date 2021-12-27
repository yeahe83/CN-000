package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yeahe83/CN-000/httpserver/metrics"
)

func main() {
	// now FROM httpserver-config/version
	// os.Setenv("VERSION", "1.0")

	metrics.Register()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// define a server
	srv := &http.Server{Addr: ":80"}
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	http.Handle("/metrics", promhttp.Handler())

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

	// 为 HTTPServer 项目添加延时 Metric
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()

	// 为 HTTPServer 添加 0-2 秒的随机延时
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))

	// 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		w.Header().Add(k, strings.Join(v, ",")) // F12 看不到
	}

	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	fmt.Fprintf(w, "VERSION = %q\n", os.Getenv("VERSION"))
	w.Header().Add("VERSION", os.Getenv("VERSION"))

	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Printf("IP = %q, Status = %q\n", r.Host, "200")

	// 转发header中的埋点信息到下级服务，以形成调用链
	req, err := http.NewRequest("GET", "http://service2.tracing.svc.cluster.local", nil)
	if err != nil {
		log.Println(err)
	}
	lowerCaseHeader := make(http.Header)
	for key, value := range r.Header {
		lowerCaseHeader[strings.ToLower(key)] = value
	}
	req.Header = lowerCaseHeader
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("get service2 fail.", err)
	}
	fmt.Fprintln(w, "==== resp from service2 ===")
	resp.Write(w)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 当访问 localhost/healthz 时，应返回200
	fmt.Fprintf(w, "200 OK")
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

