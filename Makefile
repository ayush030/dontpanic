LINT=$(GOPATH)/bin/golangci-lint

test:
	go test -v ./...

lint:
	$(LINT) run
	
.PHONY: test lint