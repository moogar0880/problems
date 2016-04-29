.PHONY: test style

### Testing
test:
	go test -v -cover -coverprofile=coverage.out

# Lint code
style:
	go vet && goimports -e -l -w .
