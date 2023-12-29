# Simple CRUD for Beginners

CRUD (Create, Read, Update, Delete) application written in Go, designed to help newcomers understand the basics of building and running a simple database-backed application.

## Makefile 

* up: Run database migrations to create tables.
* down: Roll back the database migrations.
* run: Build and run the application.
* stop: Stop and remove the PostgreSQL container.

## Create migrations

```bash
migrate create -ext sql -dir ./schema -seq init  
```

## Endpoints
* POST /users/create
* GET /users/:id   
* PUT /users/:id
* DELETE /users/:id

## Prerequisites
* Go installed on your machine.
* Docker (optional but recommended for database setup).
* run app
```
go run ./cmd/app   
```
