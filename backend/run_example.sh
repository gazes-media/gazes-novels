# Set default env vars
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=postgres
export DB_USER=postgres
export DB_PASSWORD=postgres
export PORT=5000

# Start the server
go run ./cmd/server/main.go