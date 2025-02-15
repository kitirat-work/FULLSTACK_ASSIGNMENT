reload: 
	docker compose build
	docker compose up -d

migrate-create: # make migrate-create dir=db/migrations/sql name=add_foreign_keys
	migrate create -ext sql -tz UTC -dir $(dir) $(name)

migrate-up: # make migrate-up dir=db/migrations/sql db="mysql://myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase"
	migrate -path $(dir) -database "$(db)" -verbose up

migrate-down: # make migrate-down dir=db/migrations/sql db="mysql://myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase" step=1
	migrate -path $(dir) -database "$(db)" -verbose down $(step)

migrate-force: # make migrate-force dir=db/migrations/sql db="mysql://myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase" version=20250214085440
	migrate -path $(dir) -database "$(db)" -verbose force $(version)