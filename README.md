# Go Application Template

## Directory structure
- transport: All external communications (CLI, http services and server, GRPC, ...)
- db: database related functionality (connection, migrations, seeders, ...)
- monitoring: Logger and prometheus metrics
- models: Application entities (sqlboiler generated models or handwritten structs)
- config: application configuration

## Getting Started
## Used libraries:
- SQLBoiler: for generating an ORM using a database-first approach.
- go-migrate: handling database migrations.
- Cobra: command line functionality
- Echo: Http web server
- logrus: logging tool (planning to switch to zap)
- Viper: configuration

### Dependencies
- golang-migrate command line utility: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


## CLI Built in commands
### Migration
- new
- up
- down

### models
- generate: Generate models based on configured database using sqlboiler.
