# CN-000

```
1009 模块三作业：

- 构建本地镜像。
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
- 将镜像推送至 Docker 官方镜像仓库。
- 通过 Docker 命令本地启动 httpserver。
- 通过 nsenter 进入容器查看 IP 配置。

作业需编写并提交 Dockerfile 及源代码。
提交链接：https://jinshuju.net/f/rxeJhn
截止日期：10 月 17 日晚 23:59 之前
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
