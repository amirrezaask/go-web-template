# Go Application Template

## Directory structure
- transport: All external communications (CLI, http services and server, GRPC, ...)
- db: database related functionality (connection, migrations, seeders, ...)
- monitoring: Logger and prometheus metrics
- models: Application entities (sqlboiler generated models)
- config: application configuration

## Getting Started
### Dependencies
- golang-migrate: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

