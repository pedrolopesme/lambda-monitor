GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/lambda-monitor
BINARY_UNIX=$(BINARY_NAME)_unix

build: 
	@echo "Building lambda-monitor"
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	@echo "Running lambda-monitor tests"
	$(GOTEST) -v ./...

clean: 
	@echo "Cleaning lambda-monitor"
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w */*.go

coverage:
	@echo "Running coverage via Coveralls. It expects you to have set COVERALLS_LAMBDAMONITOR_KEY env with coveralls key."
	$(GOCMD) get github.com/mattn/goveralls
	goveralls -repotoken $(COVERALLS_LAMBDAMONITOR_KEY)
