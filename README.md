# How to run

## Docker
1. Build the docker image

go build .\cmd\
 .\cmd.exe .\config\local.json

curl -d '{"CountryCode": "DE", "VatNumber": "aaa"}' -X GET http://localhost:8080/valid_vat_id
