POSTGRESQL_URL = postgres://hugo:@localhost:5432/wordsbattles?sslmode=disable
args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

run:
	go run src/main.go

watch:
	go build src/main.go
	air main.go

clean:
	rm -rf main

re:
	make clean
	make watch

build:
	go build src/main.go

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

rollback:
	migrate -path migrations/ -database ${POSTGRESQL_URL} force $(call args,error)

%:
    @:


