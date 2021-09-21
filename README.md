# How to run

## Prerequisites:
- Docker
- cURL

## Docker
1. Build the docker image:

`docker build -t vatid-validator .`

2. Run the docker container:

`docker run -p 8080:8080 -d vatid-validator`

3. Run ValidVatID endpoint with cURL:

`curl -d '{"CountryCode": "DE", "VatNumber": "123456789"}' -X GET http://localhost:8080/valid_vat_id`

## Manual building
1. Build executable:

`go build ./cmd`

2. Run executable:

`./cmd.exe config/local.json`

3. Run ValidVatID endpoint with cURL:

`curl -d '{"CountryCode": "DE", "VatNumber": "123456789"}' -X GET http://localhost:8080/valid_vat_id`

# Run tests
1. `go test ./...`