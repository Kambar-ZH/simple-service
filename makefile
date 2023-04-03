run:
	go run cmd/main.go

enter-db:
	docker exec -it postgres psql -U postgres