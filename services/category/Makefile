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
	@mockgen -destination=mock/mock_category_repository.go -package=mock -source=internal/category/port/repository.go ICategoryRepo
	@mockgen -destination=mock/mock_category_service.go -package=mock -source=internal/category/port/service.go
	@mockgen -destination=mock/mock_category_handler.go -package=mock -source=internal/category/port/handler.go