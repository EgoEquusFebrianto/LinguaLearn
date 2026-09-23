.PHONY: run build test migrate genkey clean help

help:
	@echo "Perintah:"
	@echo "  make run      - Running server"
	@echo "  make build    - Build all binaries to bin/"
	@echo "  make test     - Running test"
	@echo "  make migrate  - Running migrasi"
	@echo "  make genkey   - Generate key"
	@echo "  make clean    - Delete folder bin/"

run:
	go run cmd/server/main.go

build:
	@mkdir -p bin
	go build -o bin/server cmd/server/main.go
	go build -o bin/migrate cmd/migrate/main.go
	go build -o bin/generate_key cmd/generate_key/main.go
	@echo "Build Successful.. → bin/"

test:
	go test ./... -v

migrate:
	go run cmd/migrate/main.go

genkey:
	go run cmd/generate_key/main.go

clean:
	rm -rf bin/
	@echo "✅ Cleanup Done.."