all: serve

.PHONY: deploy

run:
	@ grep -v "#" .env | sed 's/.*/export &/' > /tmp/.env
	@ . /tmp/.env && go run main.go --db_user user --db_pwd password --db_name data-impact --db_port 27017
deploy:
	docker-compose -f ./deploy/local/docker-compose.yml up -d


teardown:
	docker-compose -f ./deploy/local/docker-compose.yml stop

release:
	@echo "Running Goreleaser..."
	goreleaser release --clean

changelog:
	@echo "Making changelog..."
	goreleaser changelog