# Go Application Template

## Directory structure
- transport: All external communications (CLI, http services and server, GRPC, ...)
- db: database related functionality (connection, migrations, seeders, ...)
- monitoring: Logger and prometheus metrics
- models: Application entities (sqlboiler generated models or handwritten structs)
- config: application configuration

## Philosophy:
This template is a bundle of the libraries and structure that I use in almost all projects, so it's tailored to my needs, 
it's not a general purpose template or framework for everyone so don't expect to find everything you need.

## Used libraries:
- SQLBoiler: for generating an ORM using a database-first approach.
- sqlx: simpler approach to database, if you just want some help in binding to structs.
- go-migrate: handling database migrations.
- Cobra: command line functionality
- Echo: Http web server
- logrus: logging tool (planning to switch to zap)
- Viper: configuration

## CLI Built in commands
### Migration
- new
- up
- down

### models
- generate: Generate models based on configured database using sqlboiler.

### Dependencies
- golang-migrate command line utility: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


