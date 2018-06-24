PACKAGES = $(shell go list ./... | grep -v '/vendor/')

#all:
#	test build docker

build:
	go build -o build/demo github.com/IaaS/services/demo
	go build -o build/demo github.com/IaaS/services/rest

docker:	
	docker build -t demo .

test:
	@echo "====> Running go test"
	@go test $(PACKAGES)

clean:
	rm -rf build

.PHONY: test
