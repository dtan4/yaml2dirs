NAME := yaml2dirs

.PHONY: build
build:
	go build -o bin/$(NAME)

.PHONY: test
test:
	go test -cover -coverprofile=coverage.txt -race -v
