# Go

test:
	go test -v .

run: up
	realize start

# Docker

up:
	docker-compose up -d

down:
	docker-compose down

genmodel:
	@read -p "Model name: " f; \
    soda generate model $${f} --skip-migration

# Migration
dbmig:
	@read -p "Migration name: " f; \
	soda generate sql $${f}
dbup:
	soda migrate up -e development

dbdown:
	soda migrate down --step 1 -e development

# Database

dbcreate:
	soda create -e development

dbdrop:
	soda drop -e development

dbreset:
	soda reset -e development
