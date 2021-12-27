export tag=1.06
build:
	echo "go build..."
	GOOS=linux GOARCH=amd64 go build -o bin/httpserver main.go

release: build
	echo "docker build..."
	docker build -t yeahe83/httpserver:${tag} .

push: release
	echo "docker push..."
	docker push yeahe83/httpserver:${tag}

run:
	echo "docker run..."
	docker run -it -p 80:80 yeahe83/httpserver:${tag}

