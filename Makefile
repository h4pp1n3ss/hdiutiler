# Define variables
BINARY_NAME := hdiutiler
OUTPUT_DIR := bin
GOFILES := ./main.go

# Create the output directory if it doesn't exist
.PHONY: prepare
prepare:
	mkdir -p $(OUTPUT_DIR)

# Build for macOS
.PHONY: build-macos
build-macos: prepare
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)_macos $(GOFILES)

# Build for both platforms
.PHONY: build
build: build-macos

# Clean up build artifacts
.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)

# Run the Go tests (if any)
.PHONY: test
test:
	go test ./...

# Install the Go tool
.PHONY: install
install: build
	cp $(OUTPUT_DIR)/$(BINARY_NAME)_macos /usr/local/bin/$(BINARY_NAME)

