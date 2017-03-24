build:
	go build -v -i -o cointicker cmd/cointicker/main.go

install:
	cd cmd && cd cointicker && \
	go install

test:
	go test ./...