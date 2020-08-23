NAME := yaml2dirs

.PHONY: build
build:
	go build -o bin/$(NAME)
