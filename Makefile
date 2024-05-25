build:
	@go build -o bin/server

run: build
	@./bin/server

run-full:
	@docker-compose --profile full down
	@docker-compose --profile full build
	@docker-compose --profile full up -d

run-db:
	@docker-compose --profile full down
	@docker-compose up -d

stop:
	@docker-compose --profile full down