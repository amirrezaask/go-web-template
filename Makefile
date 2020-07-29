install-deps:
	echo "Installing sqlboiler bins"
	GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
	GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
	echo "Installing migrate bins"
	GO111MODULE=off go get -u -d github.com/golang-migrate/migrate/cmd/migrate
	cd $(GOPATH)/src/github.com/golang-migrate/migrate/cmd/migrate
	git checkout $(TAG)
	go build -o $(GOPATH)/bin/migrate $(GOPATH)/src/github.com/golang-migrate/migrate/cmd/migrate
