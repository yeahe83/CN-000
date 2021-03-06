# CN-000

## 模块二作业

```
模块二作业
作业
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完。

接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
作业提交链接： https://jinshuju.net/f/PlZ3xg

提交截止时间：10 月 7 日晚 23:59 前
```

## 模块三作业

```
模块三作业
构建本地镜像。
编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
将镜像推送至 Docker 官方镜像仓库。
通过 Docker 命令本地启动 httpserver。
通过 nsenter 进入容器查看 IP 配置。
作业需编写并提交 Dockerfile 及源代码。

作业提交链接： https://jinshuju.net/f/rxeJhn
提交截止时间：10 月 17 日 23:59
```

```
go build -o bin/httpserver main.go
docker build -t yeahe83/httpserver:1.0 .
docker login -u yeahe83
docker push yeahe83/httpserver:1.0
docker run -it -p 8888:80 yeahe83/httpserver:1.0

docker ps|grep httpserver
docker inspect 3019|grep -i pid
nsenter -t 16086 -n ip a
curl 127.0.0.1:8888
```

## 模块八作业

```
模块八：课后作业（第一部分）
现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。

作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 kubernetes 集群，以下是你可以思考的维度

优雅启动
优雅终止
资源需求和 QoS 保证
探活
日常运维需求，日志等级
配置和代码分离
作业提交链接： https://jinshuju.net/f/OfIY3L
提交截止时间：11 月 28 日 23:59
```

```
模块八：课后作业（第二部分）
除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）

Service
Ingress
可以考虑的细节

如何确保整个应用的高可用
如何通过证书保证 httpServer 的通讯安全
结合上周模块八：课后作业（第一部分）一起提交
作业提交链接： https://jinshuju.net/f/OfIY3L
提交截止时间：11 月 28 日 23:59
```

```
# make push
# kubectl create -f httpserver-config.yaml
# kubectl create -f httpserver-svc.yaml
# kubectl create -f ingress.yaml
```

```
# curl https://cncamp.com:30621 -k
Header["X-Forwarded-Host"] = ["cncamp.com:30621"]
Header["X-Scheme"] = ["https"]
Header["User-Agent"] = ["curl/7.68.0"]
Header["X-Request-Id"] = ["aafe87b038cd73956c562dbf9789b505"]
Header["X-Forwarded-For"] = ["192.168.1.160"]
Header["X-Forwarded-Port"] = ["443"]
Header["X-Forwarded-Proto"] = ["https"]
Header["X-Forwarded-Scheme"] = ["https"]
Header["Accept"] = ["*/*"]
Header["X-Real-Ip"] = ["192.168.1.160"]
VERSION = "1.05"
```

## 模块十作业

```
1.为 HTTPServer 添加 0-2 秒的随机延时
2.为 HTTPServer 项目添加延时 Metric
3.将 HTTPServer 部署至测试集群，并完成 Prometheus 配置
4.从 Promethus 界面中查询延时指标数据

（可选）创建一个 Grafana Dashboard 展现延时分配情况
作业提交链接： https://jinshuju.net/f/z0Z07s
提交截止时间：12 月 12 日 23:59
```

```
# curl https://cncamp.com:30621/metrics -k
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 5.3001e-05
go_gc_duration_seconds{quantile="0.25"} 5.3001e-05
go_gc_duration_seconds{quantile="0.5"} 0.000193604
go_gc_duration_seconds{quantile="0.75"} 0.000297906
go_gc_duration_seconds{quantile="1"} 0.000297906
go_gc_duration_seconds_sum 0.000544511
go_gc_duration_seconds_count 3
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 11
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.17.3"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.22636e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 4.81796e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4226
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 10802
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 3.5116472734027297e-06
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 4.728496e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 2.22636e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.17792e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.555328e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 1924
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 3.825664e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 7.733248e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.639415075984815e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 12726
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 7200
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 72080
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 81920
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.216144e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.41411e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 655360
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 655360
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.4633744e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 12
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.27
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 10
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.2402688e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.63941457904e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.486155776e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 1.8446744073709552e+19
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 8
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

## 模块十二作业

```
模块十二作业
把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：

如何实现安全保证；
七层路由规则；
考虑 open tracing 的接入。
作业提交链接：  https://jinshuju.net/f/ivR6S0
提交截止时间：12 月 26 日 23:59
```

```
# kubectl create ns httpserver
# kubectl label ns httpserver istio-injection=enabled
# kubectl create -f httpserver-config.yaml -n httpserver
# kubectl create -f httpserver-svc.yaml -n httpserver
# kubectl apply -f httpserver-gateway.yaml -n httpserver
# curl --resolve httpsserver.cncamp.io:443:10.104.252.0 https://httpsserver.cncamp.io/ -k
Header["X-B3-Sampled"] = ["1"]
Header["X-Request-Id"] = ["6a1e2409-b733-9886-9ee3-db4813d0bfff"]
Header["X-B3-Traceid"] = ["6c0f8ab9b3bb57dea10e2b4be6f83dc1"]
Header["X-Envoy-Peer-Metadata-Id"] = ["router~192.168.20.90~istio-ingressgateway-55d9fb9f-wxk5v.istio-system~istio-system.svc.cluster.local"]
Header["X-Envoy-Attempt-Count"] = ["1"]
Header["Accept"] = ["*/*"]
Header["X-Envoy-Internal"] = ["true"]
Header["X-B3-Spanid"] = ["a10e2b4be6f83dc1"]
Header["User-Agent"] = ["curl/7.68.0"]
Header["X-Envoy-Decorator-Operation"] = ["httpserver.httpserver.svc.cluster.local:80/*"]
Header["X-Envoy-Peer-Metadata"] = ["ChQKDkFQUF9DT05UQUlORVJTEgIaAAoaCgpDTFVTVEVSX0lEEgwaCkt1YmVybmV0ZXMKGQoNSVNUSU9fVkVSU0lPThIIGgYxLjEyLjAKvQMKBkxBQkVMUxKyAyqvAwodCgNhcHASFhoUaXN0aW8taW5ncmVzc2dhdGV3YXkKEwoFY2hhcnQSChoIZ2F0ZXdheXMKFAoIaGVyaXRhZ2USCBoGVGlsbGVyCjYKKWluc3RhbGwub3BlcmF0b3IuaXN0aW8uaW8vb3duaW5nLXJlc291cmNlEgkaB3Vua25vd24KGQoFaXN0aW8SEBoOaW5ncmVzc2dhdGV3YXkKGQoMaXN0aW8uaW8vcmV2EgkaB2RlZmF1bHQKMAobb3BlcmF0b3IuaXN0aW8uaW8vY29tcG9uZW50EhEaD0luZ3Jlc3NHYXRld2F5cwofChFwb2QtdGVtcGxhdGUtaGFzaBIKGgg1NWQ5ZmI5ZgoSCgdyZWxlYXNlEgcaBWlzdGlvCjkKH3NlcnZpY2UuaXN0aW8uaW8vY2Fub25pY2FsLW5hbWUSFhoUaXN0aW8taW5ncmVzc2dhdGV3YXkKLwojc2VydmljZS5pc3Rpby5pby9jYW5vbmljYWwtcmV2aXNpb24SCBoGbGF0ZXN0CiIKF3NpZGVjYXIuaXN0aW8uaW8vaW5qZWN0EgcaBWZhbHNlChoKB01FU0hfSUQSDxoNY2x1c3Rlci5sb2NhbAotCgROQU1FEiUaI2lzdGlvLWluZ3Jlc3NnYXRld2F5LTU1ZDlmYjlmLXd4azV2ChsKCU5BTUVTUEFDRRIOGgxpc3Rpby1zeXN0ZW0KXQoFT1dORVISVBpSa3ViZXJuZXRlczovL2FwaXMvYXBwcy92MS9uYW1lc3BhY2VzL2lzdGlvLXN5c3RlbS9kZXBsb3ltZW50cy9pc3Rpby1pbmdyZXNzZ2F0ZXdheQoXChFQTEFURk9STV9NRVRBREFUQRICKgAKJwoNV09SS0xPQURfTkFNRRIWGhRpc3Rpby1pbmdyZXNzZ2F0ZXdheQ=="]
Header["X-Forwarded-For"] = ["192.168.1.160"]
Header["X-Forwarded-Proto"] = ["https"]
VERSION = "1.05"
==== resp from service2 ===
HTTP/1.1 200 OK
Content-Length: 647
Content-Type: text/plain; charset=utf-8
Date: Mon, 27 Dec 2021 11:40:42 GMT
Server: envoy
X-Envoy-Upstream-Service-Time: 194

===================Details of the http request header:============
X-Forwarded-For=[192.168.1.160]
X-Forwarded-Client-Cert=[By=spiffe://cluster.local/ns/tracing/sa/default;Hash=d10ed52cf9d76682ae595e3093b966ca43f1b048e3f2f5e89656a1da662222d3;Subject="";URI=spiffe://cluster.local/ns/httpserver/sa/default]
Accept=[*/*]
X-Envoy-Attempt-Count=[1]
X-Request-Id=[6a1e2409-b733-9886-9ee3-db4813d0bfff]
User-Agent=[Go-http-client/1.1,curl/7.68.0]
Accept-Encoding=[gzip]
X-Envoy-Internal=[true]
X-B3-Parentspanid=[1e2dcdb2f078c249]
X-Forwarded-Proto=[https]
X-B3-Spanid=[8e4b361b13b15786]
X-B3-Sampled=[1]
X-B3-Traceid=[6c0f8ab9b3bb57dea10e2b4be6f83dc1]

