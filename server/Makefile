ContainerNamePostgres = postgreslocalaom
export PostgresUser = root
export PostgresPassword = password 
export PostgresPort = 3100
export DBName= go-chat
export DBHost= 127.0.0.1
postgresinit:
	docker run --name $(ContainerNamePostgres) -p $(PostgresPort):5432 -e POSTGRES_USER=$(PostgresUser) -e POSTGRES_PASSWORD=$(PostgresPassword) -d postgres:15-alpine

postgresdb:
	docker exec -it $(ContainerNamePostgres) psql



createdb:
	docker exec -it $(ContainerNamePostgres) createdb --username=$(PostgresUser) --owner=$(PostgresUser) $(DBName)

dropdb:
	docker exec -it $(ContainerNamePostgres) dropdb $(DBName)

.PHONY: postgresinit postgresdb createdb dropdb