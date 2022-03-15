## Install golangci-lint locally
lint.install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	golangci-lint --version

## Run golangci-lint installed locally
lint.run:
	golangci-lint run -c .golangci.yml -v --fix


