FROM golang:1.11

COPY . /app
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vatid-validator

EXPOSE 8080
ENTRYPOINT ["./vatid-validator", "/app/config/local.json"]