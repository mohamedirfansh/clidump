APP_NAME=clidump
CMD_DIR=./cmd/$(APP_NAME)

.PHONY: dev build test clean

dev:
	go run $(CMD_DIR) -t "list all files in current directory including hidden files"

build:
	go build -o bin/$(APP_NAME) $(CMD_DIR)

test:
	go test ./...

clean:
	rm -rf bin
