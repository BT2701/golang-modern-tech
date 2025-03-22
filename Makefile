# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=myapp
BINARY_UNIX=$(BINARY_NAME)_unix

# All target
all: test build

# Build the project
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run tests
test: 
	$(GOTEST) -v ./...

# Clean build files
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Install dependencies
deps:
	$(GOGET) -u ./...

# Cross compilation for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

start-roadmap:
	go run weekly_roadmap/cmd/main.go

start-project:
	go run modern-tech/mini_project/cmd/main.go