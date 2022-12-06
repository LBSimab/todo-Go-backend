createdb: 
	docker exec -it tasks createdb --username=icarus --owner=icarus tasksdb
dropdb: 
	docker exec -it tasks dropdb  tasksdb --username=icarus+
test:
	go test -v -cover ./...

sqlc:
	sqlc generate


migrateup:
		migrate -path db/migration -database "postgresql://icarus:secret@localhost:5432/tasksdb?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://icarus:secret@localhost:5432/tasksdb?sslmode=disable" -verbose down










.PHONY: createdb dropdb test sqlc migaretup migratedown