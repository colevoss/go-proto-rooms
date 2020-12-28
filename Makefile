gen:
	bash scripts/protoc-gen.sh

build:
	go build ./cmd/server/main.go