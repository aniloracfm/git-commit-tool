BINARY_NAME=commit-tool
CLI_PATH=./cmd/cli/main.go

.PHONY: build

## help - Comandos disponíveis no Makefile.
help:
	@echo "Comandos disponíveis:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':'

## setup - Configura o ambiente de desenvolvimento.
setup:
	@chmod +x scripts/config-automation.sh
	@bash scripts/config-automation.sh

## run - Executa o projeto - apenas para fins de teste.
run:
	@go run $(CLI_PATH)

## build - Compila o projeto.
build:
	@mkdir -p bin
	@go build -o bin/$(BINARY_NAME) $(CLI_PATH)

## test - Executa os testes unitários de todos os pacotes.
test:
	@go test -v ./...

## clean - Remove os testes unitários de todos os pacotes.
clean:
	@rm -f bin/$(BINARY_NAME)	


