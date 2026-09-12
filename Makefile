.PHONY: help build test lint fmt fmt-check vet check run

help:
	@echo "Доступные цели:"
	@echo "  make build      - сборка проекта"
	@echo "  make test       - запуск тестов"
	@echo "  make lint       - линтер"
	@echo "  make fmt        - форматирование кода"
	@echo "  make fmt-check  - проверка форматирования"
	@echo "  make vet        - go vet"
	@echo "  make check      - fmt-check + vet + build + lint + test"
	@echo "  make run        - запуск приложения"

build:
	go build ./...

test:
	go test ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .

fmt-check:
	@out=$$(gofmt -l .); \
	if [ -n "$$out" ]; then echo "Требуется gofmt:"; echo "$$out"; exit 1; fi

vet:
	go vet ./...

check: fmt-check vet build lint test

run:
	go run ./cmd/todo
