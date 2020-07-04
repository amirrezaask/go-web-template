models-generate:
	go generate
migration-generate:
	migrate create -ext sql -dir db/migrations -seq $(name) 
