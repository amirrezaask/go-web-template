# Go Application Template
## Dependencies
### SQLBoiler
    - GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
    - GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
### go-migrate
    - GO111MODULE=off go get -u -d github.com/golang-migrate/migrate/cmd/migrate
    - cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
    - git checkout $TAG
    - go build -o $GOPATH/bin/migrate $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
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
- SQLBoiler: for generating an ORM using a database-first approach. (recommended only for applications with a lot of tables)
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
### swag
- generate: generate  models for http handlers based on swagger API spec
### GRPC:
- generate: generate GRPC go files based on given proto file
### Dependencies
- golang-migrate command line utility: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


