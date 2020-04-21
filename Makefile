# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOROOT:=$$(go env GOROOT)
BASE_CERTPATH=/src/crypto/tls/generate_cert.go
CERTPATH=$(GOROOT)$(BASE_CERTPATH)
BINARY_NAME=eospke
BINARY_UNIX=$(BINARY_NAME)_unix
BUILD_SWAGGER=swag init

all: test build
build: 
	$(BUILD_SWAGGER)
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(BUILD_SWAGGER)
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)
devcert:	
	$(GOCMD) run $(CERTPATH) --host localhost
swagger:
	$(BUILD_SWAGGER)
deps:
	$(GOCMD) mod download
	$(GOCMD) get -u github.com/swaggo/swag/cmd/swag