.SILENT:

run: 
	docker-compose up --build
migrate:
	docker run -v ./schema:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:ShatAlex@localhost:5433/postgres?sslmode=disable up