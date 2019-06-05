.PHONY: build clean deploy

build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/echo functions/post/echo/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello_page functions/get/hello_page/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/post_email functions/post/email/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get_email functions/get/email/main.go

clean:
	rm -rf ./bin ./vendor 

deploy: clean build
	sls deploy --verbose
