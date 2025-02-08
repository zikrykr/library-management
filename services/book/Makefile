MIGRATE := migrate
MIGRATIONS_DIR := config/db/migration

migrate-up:
	@$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate-down:
	@$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) down $(STEP)

create-migration:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

mockgen:
	@go install github.com/golang/mock/mockgen
	@mockgen -destination=mock/mock_book_repository.go -package=mock -source=internal/book/port/repository.go IBookRepo
	@mockgen -destination=mock/mock_book_service.go -package=mock -source=internal/book/port/service.go
	@mockgen -destination=mock/mock_book_handler.go -package=mock -source=internal/book/port/handler.go