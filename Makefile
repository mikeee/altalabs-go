ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -count=1 \
			-covermode=atomic \
			-coverprofile=coverage.out \
			-v \
			./...

.PHONY: test-e2e
test-e2e:
	go test github.com/mikeee/altalabs-go/test/e2e/ \
			-tags e2e \
			-count=1 \
			-covermode=atomic \
			-coverprofile=coverage-e2e.out \
			-v

.PHONY: run-example
run-example:
	go run examples/basic/main.go