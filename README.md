# Rezerve: Service Microservice 

# Setup
## Docker

To set up the project using Docker, follow these steps:

1. Run `cp -n .env.dist .env`.
2. Execute `make init` to initialize the project.
3. Start the development server by running `make start`.

Now, the application should be accessible at http://127.0.0.1:9030. You can test its availability by [pinging it](http://ttp://127.0.0.1:9030/v1/health/ping).

## Available Commands

Here are some useful commands for managing the project:

- `make tf-plan`: Check the current Terraform state.
- `make tests`: Run all tests, output coverage, and save coverage to `coverage.out`.
- `make coverage`: View the coverage report in your local browser. Ensure you run `make tests` first to generate the coverage report.

## Stack

Here are the main technologies and tools used in this project:

1. `gofiber.io`: GoLang dependency injection toolkit.
2. `golang 1.22`
3. `terraform`: Used for DynamoDB schema definitions.