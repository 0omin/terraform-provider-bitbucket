.PHONY: build test vet fmt

build:
	go build -o terraform-provider-bitbucket .

test:
	go test ./... -race

vet:
	go vet ./...

fmt:
	gofmt -l -w .
