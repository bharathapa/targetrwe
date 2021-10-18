fmt:
	go fmt -n ./...
lint:
	golangci-lint run
run:
	go run cmd/main.go