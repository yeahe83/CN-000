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
