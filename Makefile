help:
	@echo "install 		Install all dependencies declared on Gokpg.toml using deps"
	@echo "test 			Run tests and generates coverage data/assets"

install:
	@dep ensure -v

test:
	@mkdir -p coverage
	@go test -v ./... -coverprofile=c.out
	@go tool cover -html=c.out -o coverage/index.html