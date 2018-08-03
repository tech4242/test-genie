# Go parameters
GOCMD=go
GOFMT=gofmt
GOBUILD=go build
GOCLEAN=go clean
GOGET=go get
# Go paths
PACKAGE=server
BINARY_PATH=./bin
BINARY=$(BINARY_PATH)/$(PACKAGE)
SRC=$(wildcard src/*.go)

$(PACKAGE):$(SRC)
	$(GOFMT) -w -d $(SRC)
	$(GOBUILD) -v -o $(BINARY) $(SRC)

all: $(PACKAGE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY)

run:
	$(BINARY) -redirect-url=$(REDIRECTURL)

deps:
	$(GOGET) github.com/gorilla/mux
