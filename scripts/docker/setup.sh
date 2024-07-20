#!/bin/bash

set -euo pipefail

DIR=$(dirname "$0")

echo " > Configuring settings"
docker-compose run --rm app sh scripts/docker/config_settings.sh

echo " > Update packages"
docker-compose run --rm app go mod vendor

echo " > Bring containers up"
docker-compose up -d app localstack

echo " > Terraform init"
docker-compose run --rm terraform init

echo " > Terraform plan with output"
docker-compose run --rm terraform plan -out=terraform-plan

echo " > Terraform apply previously create output "
docker-compose run --rm terraform apply terraform-plan

echo " > DynamoDB fixtures "
docker-compose exec app go run scripts/fixture/main.go
