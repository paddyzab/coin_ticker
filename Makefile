build:
	go build -v -i ./cmd/cointicker

install:
	go install ./cmd/cointicker

setup:
	go get github.com/urfave/cli github.com/logrusorgru/aurora

test:
	go test ./...

test-setup:
	go get github.com/stretchr/testify/assert
