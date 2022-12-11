sqlc: 
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

migrateup: 
	docker run -v "%cd%/db/migration:/migrations" --network host migrate/migrate -path=/migrations/ -database postgresql://root:1@localhost:5432/simple_bank?sslmode=disable up 2

migratedown: 
	docker run -v "%cd%/db/migration:/migrations" --network host migrate/migrate -path=/migrations/ -database postgresql://root:1@localhost:5432/simple_bank?sslmode=disable down 2 



.PHONY: sqlc migrateup