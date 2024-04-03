all: serve

.PHONY: deploy

run:
	@ grep -v "#" .env | sed 's/.*/export &/' > /tmp/.env
	@ . /tmp/.env && go run main.go --db_user user --db_pwd password --db_name data-impact --db_port 27017
deploy:
	docker-compose -f ./deploy/local/docker-compose.yml up -d


teardown:
	docker-compose -f ./deploy/local/docker-compose.yml stop

CURRENT_VERSION := $(shell git describe --tags --abbrev=0)

BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)

# Increment version based on branch name prefix
ifeq ($(findstring feature:,$(BRANCH_NAME)),feature:)
  NEW_VERSION := $(shell echo $(CURRENT_VERSION) | awk -F. '{$$NF = $$NF + 1;} 1' | sed 's/^v//')
else ifeq ($(findstring bug:,$(BRANCH_NAME)),bug:)
  NEW_VERSION := $(shell echo $(CURRENT_VERSION) | awk -F. '{$$(NF-1) = $$(NF-1) + 1;} 1' | sed 's/^v//')
else
  $(error Branch name does not start with 'feature:' or 'bug:')
endif

.PHONY: release
release: tag goreleaser

.PHONY: tag
tag:
ifeq ($(shell git tag -l $(NEW_VERSION)),)
	@echo "Tagging new version: v$(NEW_VERSION)"
	@git tag -a v$(NEW_VERSION) -m "Version $(NEW_VERSION) release"
	@git push origin v$(NEW_VERSION)
else
	@echo "Tag already exists: v$(NEW_VERSION)"
endif

.PHONY: goreleaser
goreleaser:
	@echo "Running Goreleaser..."
	@goreleaser release