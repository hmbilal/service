.PHONY: init start bye ssh tests coverage lint lint-fix mocks fixture tf-init tf-plan-out tf-apply tf-plan setup-test-db migrate project-setup


init:
	./scripts/docker/init.sh

start:
	docker-compose up -d

bye:
	docker-compose down

ssh:
	docker-compose exec app sh

tests:
	docker-compose exec app go test -failfast -race -mod vendor ./... -v -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

lint:
	docker-compose exec app golangci-lint run --deadline=5m

lint-fix:
	docker-compose exec app golangci-lint run --modules-download-mode vendor --deadline=5m --fix

mocks:
	docker-compose exec app sh scripts/generate-mocks.sh

fixture:
	docker-compose exec app go run scripts/fixture/main.go

tf-init:
	docker-compose run --rm terraform init

tf-plan-out:
	docker-compose run --rm terraform plan -out=terraform-plan

tf-apply:
	docker-composee run --rm terraform apply terraform-plan

tf-plan:
	docker-compose run --rm terraform plan

setup-test-db:
	docker-compose up -d localstack-test
	docker-compose run --rm terraform plan -out=terraform-test-plan -var="db_url=http://localstack-test:4566"
	docker-compose run --rm terraform apply terraform-test-plan

migrate:
	make tf-plan-out && make tf-apply

project-setup:
	./scripts/setup_project.sh $(PLACEHOLDER) $(NAME)