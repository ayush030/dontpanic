LINT=$(GOPATH)/bin/golangci-lint

test:
	go test -v ./...

lint:
	$(LINT) run

build-plugin:
	mkdir -p build
	go build -buildmode=plugin -o build/dontpanic cmd/main.go

build-standalone:
	mkdir -p build
	go build -o build/dontpanic cmd/main.go
clean: 
	rm -rf build/*

.PHONY: test lint clean build-plugin build-standalone
