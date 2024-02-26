SERVICE_NAME=txt-thumbnailer

# Сборка сервиса
.PHONY: build
build:
	go build -o bin/$(SERVICE_NAME) -ldflags=$(LD_FLAGS) $(PWD)/cmd/$(SERVICE_NAME)
	echo "build successfully"

test:
	go test ./...

example:
	go run cmd/txt-thumbnailer/main.go  convert examples/txt/lorem.txt  --padding-left=100 --padding-top=100

server:
	go run cmd/txt-thumbnailer/main.go server