# Define your commands as targets
fmt:
	go fmt ./...

build:
	go build ./...

test:
	go test ./...

run:
	go run cmd/api/main.go
