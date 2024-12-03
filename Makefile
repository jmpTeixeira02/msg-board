default: build

build:
	go build -o ./build/server cmd/main.go

generate:
	mkdir -p daemon/api/generated
	oapi-codegen -config daemon/api/oapi-config.yml daemon/api/swagger.yml

deps:
	go mod download
	go mod tidy

clean:
	rm -rf ./tmp/

.PHONY: build

