.DEFAULT_GOAL := build
.PHONY: fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

run:
	go run ./main/main.go ./archivo.txt

clean:
	rm -f ./main/main
