include .env
DB_CONNECTION = mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):3306)/$(DB_NAME)?multiStatements=true

run:
	go run main.go

test:
	go test -v ./...

migrate-up:
	migrate -database "$(DB_CONNECTION)" \
	-path db/migrations up

migrate-down:
	migrate -database "$(DB_CONNECTION)" \
	-path db/migrations down

make-migration:
ifdef name
	@migrate create -dir db/migrations -ext sql -seq $(name)
else
	@echo "'name' is not provide; e.g. make-migration name=user"
endif
