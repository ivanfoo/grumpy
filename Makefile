SHELL = /bin/bash
PROJECT = grumpy

dependencies:
	go get -v -t ./...

artifacts:
	@mkdir artifacts \
	&& go build -o artifacts/dummy \

clean:
	@rm -rf artifacts/

rebuild: clean artifacts
