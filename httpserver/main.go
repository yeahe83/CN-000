/*
9.25课后作业
内容：编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把1都做完

1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回200
提交链接🔗：https://jinshuju.net/f/PlZ3xg
截止时间：10月7日晚23:59前
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	os.Setenv("VERSION", "1.0")

	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)

	log.Fatal(http.ListenAndServe(":80", nil))
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
