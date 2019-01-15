# test all
.PHONY: test
test: tidy
	go test ./...

# tidy dependencies
.PHONY: tidy
tidy: install
	go mod tidy

# install all
.PHONY: install
install: lint
	go install ./...

# download golint and lint
BIN_DIR := $(GOPATH)/bin
GOLINT := $(BIN_DIR)/golint
$(GOLINT):
	go get -u golang.org/x/lint/golint
	
.PHONY: lint
lint: $(GOLINT) fmt
	golint ./...

# format
.PHONY: fmt
fmt:
	go fmt ./...


