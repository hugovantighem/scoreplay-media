# TLDR
Clone the repository

run the application
```
make copy-env-file
make start-db
make run-scratch
```

run all the tests with coverage
```
make start-db
make test-all-coverage
```


# Run the application

`make start-db` then `make run-scratch`


NOTE: the application requires a running mongoDB that can be run using: `make start-db` (docker required).

NOTE: the application requires a `.env` file for variable configuration. You can set you own configuration based on `.env-example` file. `make copy-env-file` creates a `.env` file compatible with `make start-db` command.

NOTE: requires docker

curl commands for testing on local server are provided into `client` directory.

# Development
The `Makefile` contains targets to run individual target, or any target (`*-scratch`) along with all previously required targets as if you just cloned the repository.

`make install-deps` for generation tools

`make generate` for api and mock generation

`make build` to build binary

`make run` to run the application. 


# Tests
run unit tests
```
make test-unit-scratch
```

run integration tests
```
make start-db
make test-intgration-scratch
```

NOTE: integration tests use mongodb from docker, and run a local application

## Coverage
test coverage can be computed including only unit tests or all the tests:

```
make start-db
make test-unit-coverage
```

```
make start-db
make test-all-coverage
```