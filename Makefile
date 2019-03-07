help:
	@echo "install 		Install all dependencies declared on Gokpg.toml using deps"
	@echo "test 			Run tests and generates coverage data/assets"

install:
	@dep ensure -v

test:
	@mkdir -p coverage
	@go test -v ./... -coverprofile=coverage/coverage.out
	@go tool cover -html=coverage/coverage.out -o coverage/index.htm