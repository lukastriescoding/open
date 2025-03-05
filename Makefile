BIN_DIR=bin
BIN_NAME=open
BIN_PATH=$(BIN_DIR)/$(BIN_NAME)

.PHONY: build open clean

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_PATH)

build-windows
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o open.exe

open: build
	@$(BIN_PATH)

clean:
	rm -rf $(BIN_DIR)

clean-home:
	rm -rf $(HOME)/.open
