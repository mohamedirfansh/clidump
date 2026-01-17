APP_NAME=clidump
CMD_DIR=./cmd/$(APP_NAME) -t "list all files in this directory"

.PHONY: dev build test clean

dev:
	go run $(CMD_DIR)

build:
	go build -o bin/$(APP_NAME) $(CMD_DIR)

test:
	go test ./...

clean:
	rm -rf bin
