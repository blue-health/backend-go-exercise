NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

SERVICE_NAME=backend-go-exercise-srv
 
.PHONY: all format build
all: format build

GO_EXERCISE_DATABASE_DSN := postgres://example:qwerty123@localhost/example?sslmode=disable

# App

run:
	@echo "$(OK_COLOR)==> Running $(SERVICE_NAME)...$(NO_COLOR)"
	@GO_EXERCISE_DATABASE_DSN=$(GO_EXERCISE_DATABASE_DSN) go run main.go

build: test
	@echo "$(OK_COLOR)==> Building $(SERVICE_NAME)...$(NO_COLOR)"
	@go build -o bin/$(SERVICE_NAME) main.go
 
compile:
	@echo "$(OK_COLOR)==> Compiling $(SERVICE_NAME) for Linux x86-64...$(NO_COLOR)"
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/$(SERVICE_NAME)-linux-amd64 main.go

test: lint
	@echo "$(OK_COLOR)==> Testing $(SERVICE_NAME)...$(NO_COLOR)"
	@go test

lint: 
	@echo "$(OK_COLOR)==> Linting $(SERVICE_NAME)...$(NO_COLOR)"
	@golangci-lint run

format:
	@echo "$(OK_COLOR)==> Formatting $(SERVICE_NAME)...$(NO_COLOR)"
	@go fmt

# Database

database-up:
	@echo "$(OK_COLOR)==> Running docker compose for $(SERVICE_NAME)...$(NO_COLOR)"
	@docker-compose up

database-down:
	@echo "$(OK_COLOR)==> Running docker compose for $(SERVICE_NAME)...$(NO_COLOR)"
	@docker-compose down
