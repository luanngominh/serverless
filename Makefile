.PHONY: build clean deploy

build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/echo functions/post/echo/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello_page functions/get/hello_page/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
