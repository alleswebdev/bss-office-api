.PHONY: build
build:
	go build cmd/omp-bss-office-api/main.go

.PHONY: test
test:
	go test -v ./...