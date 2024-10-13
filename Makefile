check: lint test

fmt:
	go fmt ./...

fmtfix:
	golangci-lint run --fix -E gofmt,gofumpt,goimports

lint:
	golangci-lint run

test:
	go test ./...

bench:
	go test -bench=. ./...

run:
	go run main.go