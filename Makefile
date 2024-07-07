tidy:
	@go mod tidy

build:
	@go build -o bin/gobank

run:
	@go run main.go

# run: build
# 	@./bin/gobank

test:
	@go test -v ./..