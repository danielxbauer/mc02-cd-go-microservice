# MC02 - Continuous Delivery - Go Microservice
Following tutorial from: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

## Database
Create the database by:
* `docker pull postgres`
* `docker run -it --name cd-postgres -p 5432:5432 -e POSTGRES_PASSWORD=<password> -d postgres`
* execute `data/init.sql`

## Environment Variables
Following environment variables need to be set:
* See `src/.env` file

## Build
Build the project by installing `Go` and executing `go build -v`.

## Test
Test the project by executing `go test -v`