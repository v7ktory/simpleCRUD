up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

run:
	docker run --name=yourname -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres  
	
stop: 
	docker stop yourname