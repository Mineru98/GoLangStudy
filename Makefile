hello:
	@echo "Hello World"

run:
	@go run src/example/main.go

build:
	@go build -o bin/main.exe src/example/main.go
	@bin/main.exe

test:
	@go test -v

all: hello run build test

cc:
	@echo "Cross Complie"
	set GOOS=darwin& set GOARCH=386& go build -o bin/main-darwin-386.exe src/example/main.go
	set GOOS=darwin& set GOARCH=arm64& go build -o bin/main-darwin-arm64.exe src/example/main.go
	set GOOS=darwin& set GOARCH=arm& go build -o bin/main-darwin-arm.exe src/example/main.go
	set GOOS=linux& set GOARCH=arm64& go build -o bin/main-linux-amd64.exe src/example/main.go
	set GOOS=linux& set GOARCH=arm& go build -o bin/main-linux-amd64.exe src/example/main.go
	set GOOS=windows& set GOARCH=arm64& go build -o bin/main-windows-arm.exe src/example/main.go
	set GOOS=windows& set GOARCH=arm& go build -o bin/main-windows-arm.exe src/example/main.go

cclinux:
	@echo "Cross Complie"
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386.exe src/example/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/main-darwin-arm64.exe src/example/main.go
	GOOS=darwin GOARCH=arm go build -o bin/main-darwin-arm.exe src/example/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-amd64.exe src/example/main.go
	GOOS=linux GOARCH=arm go build -o bin/main-linux-amd.exe src/example/main.go
	GOOS=windows GOARCH=arm64 go build -o bin/main-windows-arm64.exe src/example/main.go
	GOOS=windows GOARCH=arm go build -o bin/main-windows-arm.exe src/example/main.go