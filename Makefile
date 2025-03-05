BIN_DIR=bin
BIN_NAME=open
BIN_PATH=$(BIN_DIR)/$(BIN_NAME)

.PHONY: build open clean

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_PATH)

open: build
	@$(BIN_PATH)

clean:
	rm -rf $(BIN_DIR)

clean-home:
	rm -rf $(HOME)/.open
