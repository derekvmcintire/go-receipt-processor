# Define your commands as targets
fmt:
	go fmt ./...

build:
	go build ./...

test:
	go test ./...

run:
	go run cmd/api/main.go

docker-build:
	docker build -t receipt-processor .

docker-run:
	docker run -p 8080:8080 receipt-processor

