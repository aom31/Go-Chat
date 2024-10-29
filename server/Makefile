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


namefile=add_users_table
dbConnSchemaOnboardingDev = 'postgresql://root:password@localhost:3100/go-chat?sslmode=disable'
migrate-create:
	migrate create -ext sql -dir db/migrations $(namefile)

migrate-up:
	migrate -database $(dbConnSchemaOnboardingDev) -path db/migrations -verbose up  

migrate-down:
	migrate -database $(dbConnSchemaOnboardingDev) -path db/migrations -verbose down  


version=20241029092242
force-migrate-version:
	migrate -path db/migrations -database $(dbConnSchemaOnboardingDev) force $(version)
