PACKAGE=github.com/moogar0880/problems

.PHONY: all test coverage style

all: test

build:
	go build -v $(PACKAGE)

### Testing
test: build
	go test -v $(PACKAGE)

coverage:
	go test -v -cover -coverprofile=coverage.out $(PACKAGE)

### Style and Linting
lint:
	go vet $(PACKAGE) && goimports .

# modify source code if style offences are found
style:
	go vet $(PACKAGE) && goimports -e -l -w .
