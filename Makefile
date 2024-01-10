test:
	go test ./...

example:
	go run cmd/txt-thumbnailer/main.go  convert examples/txt/lorem.txt  --padding-left=100 --padding-top=100
