.PHONY: test fmt

# Форматирование кода
fmt:
	go fmt ./...

# Запуск тестов
test:
	go test -v ./...

# Запуск тестов с покрытием
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Очистка временных файлов
clean:
	rm -f coverage.out

# Установка зависимостей
deps:
	go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Помощь
help:
	@echo "Доступные команды:"
	@echo "  make fmt       - Форматирование кода"
	@echo "  make test      - Запуск тестов"
	@echo "  make test-cover- Запуск тестов с проверкой покрытия"
	@echo "  make deps      - Установка зависимостей"
	@echo "  make clean     - Очистка временных файлов"
