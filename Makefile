###############################################################################
# GNU Make Variables
###############################################################################
export MAKEFLAGS += --warn-undefined-variables
export SHELL := bash
export .SHELLFLAGS := -eu -o pipefail -c
export .DEFAULT_GOAL := all
.DELETE_ON_ERROR:
.SUFFIXES:

###############################################################################
# Location Variables
###############################################################################
export PROJECT_ROOT=$(shell pwd)

###############################################################################
# Go Variables
###############################################################################
GO_MODULE_NAME=github.com/moogar0880/problems
GO_COVERAGE_FILE=cover.out
GO_TEST_OPTS=-coverprofile $(GO_COVERAGE_FILE)
GO_TEST_PKGS=./...

.PHONY: clean
clean:
	@rm $(GO_COVERAGE_FILE)

.PHONY: godoc
godoc:
	docker run \
		--rm \
		--detach \
		--entrypoint bash \
		--expose "6060" \
		--name "godoc" \
		--publish "6060:6060" \
		--volume $(PROJECT_ROOT):/go/src/$(GO_MODULE_NAME) \
		--workdir /go/src/$(GO_MODULE_NAME) \
		golang:latest \
		-c "go install golang.org/x/tools/cmd/godoc@latest && godoc -http=:6060" || true
	@open http://localhost:6060

### Testing
.PHONY: test
test:
	@go test $(GO_TEST_OPTS) $(GO_TEST_PKGS)

.PHONY: test/coverage
test/coverage: test
	@go tool cover -html=cover.out
