.PHONY: build run run-api doc validate spec clean web help

all: validate clean build

validate:
	swagger validate ./api/swagger.yml

spec:
	swagger generate spec -o ./api/swagger-gen.yml

build:
	swagger -q generate server -A youtube-api -f api/swagger.yml -s internal/apis -m internal/models
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v ./cmd/youtube-api-server


api: validate clean build
	./youtube-api-server --port=8080 --host=0.0.0.0

run: all
	./youtube-api-server --port=8099 --host=0.0.0.0

doc:
	swagger validate ./api/swagger.yml
	swagger serve api/swagger.yml --no-open --host=0.0.0.0 --port=9090 --base-path=/

clean:
	rm -rf youtube-api-server
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make validate: OpenAPI validation"
	@echo "make spec: OpenAPI Spec"
	@echo "make clean: remove object files and cached files"
	@echo "make build: Generate Server and Client API"
	@echo "make doc: Serve the Doc UI"
	@echo "make web: Build svelte static"
	@echo "make run-api: Build api only and serve"
	@echo "make run: Build everything and serve"
