createdb:
	docker exec -it postgres12 createdb --username=root --owner=root amex_example
dropdb:
	docker exec -it postgres12 dropdb amex_example
migrateup:
	migrate -path repository/migrations -database "postgresql://root:secret@localhost:5432/amex_example?sslmode=disable" -verbose up
migratedown:
	migrate -path repository/migrations -database "postgresql://root:secret@localhost:5432/amex_example?sslmode=disable" -verbose down

.PHONY: createdb, dropdb