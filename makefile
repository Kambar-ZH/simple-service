run:
	go run cmd/main.go

enter-db:
	docker exec -it postgres psql -d diploma_project -U postgres

swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/main.go -o ./docs