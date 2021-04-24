NAME:=gel
EXAMPLES:=$(shell ls -1 examples/)

all: test build

test:
	go test -v --short ./...

lint:
	golangci-lint run

.PHONY: examples
examples:
	@for i in $(EXAMPLES); do \
		DIR=examples/$$i; \
		echo "\033[1;037mRunning example: $$i...\033[0m"; \
		(cd $$DIR && ./run.sh); \
		echo ""; \
	done
