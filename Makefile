up: up-db up-app ## Запускает docker-compose

up-db: ## Запускает docker-compose для базы данных
	@docker-compose up -d db

up-app: ## Ждет готовности базы данных и запускает остальные контейнеры
	@echo "Ожидание запуска базы данных..."
	@sleep 10 # Задержка в 10 секунд для ожидания готовности базы данных
	@docker-compose up --build swagger app

test:
	go test -v

local:	
	go run ./cmd/main.go
