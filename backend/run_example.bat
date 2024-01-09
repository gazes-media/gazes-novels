:: Set Environment Variables for Windows

set DB_HOST=localhost
set DB_PORT=5432
set DB_NAME=postgres
set DB_USER=postgres
set DB_PASSWORD=postgres
set PORT=5000

:: Run the server in development mode
go run ./cmd/server/main.go

