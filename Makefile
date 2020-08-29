
all: build run

build:
	@GOOS=linux GOARCH=386 go build -o bin/linux-386/monitorer-linux-386 main.go monitore.go
	@GOOS=windows GOARCH=386 go build -o bin/windows-386/monitorer-windows-386.exe main.go monitore.go

run:
	@go run main.go monitore.go