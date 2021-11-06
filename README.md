# phone-list-be

# Architecture
based on 4 layer
* Model
* Repository
* Service
* Handler

## Features
- [x] containerized using `docker` and `docker-compose`
- [x] API Documentation using `swagger` (auto generated)
- [x] `pagination`
- [x] Middlewares `CORS`, `Logger`, `Recover`
- [x] Unit tests
- [ ] Benchmark
- [ ] Code Docs

## Requirements
* using docker
    * docker
    * docker-compose
* without docker
    * golang (tested on `go1.17.2`)

## Run using docker
* run `docker-compose up -d`
* open http://localhost:5001/swagger/index.html

## Run without docker
* run `go run main.go`
* open http://localhost:5001/swagger/index.html

## Run tests
* run `make test`
