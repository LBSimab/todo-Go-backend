createdb: 
	docker exec -it tasks createdb --username=icarus --owner=icarus tasksdb
dropdb: 
	docker exec -it tasks dropdb  tasksdb --username=icarus+
test:
	go test -v -cover ./...

sqlc:
	sqlc generate












.PHONY: createdb dropdb test sqlc