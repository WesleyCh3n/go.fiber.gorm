
host     := localhost
user     := wesley
db       := testdb
password := 970530
port     := 5432
sslmode  := disable

postgres:
	docker run --name postgres_ctr -p $(port):$(port) -e POSTGRES_USER=$(user) -e POSTGRES_PASSWORD=$(password) -d postgres

createdb:
	docker exec -it postgres_ctr createdb --username=$(user) --owner=$(user) $(db)

dropdb:
	docker exec -it postgres_ctr dropdb $(db)

.PHONY: postgres createdb dropdb
