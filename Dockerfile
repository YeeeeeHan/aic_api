# To build the Docker image, run from project root:
# docker build -t <REPOSITORY_NAME>:<TAG_NAME> <PATH_TO_DOCKERFILE>
# e.g.: docker build -t aic_api:latest ..

# Build stage
FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
RUN apk --no-cache add tzdata
COPY --from=builder /app/main .
COPY --from=builder /app/common ./common
COPY .env .

EXPOSE 8888
# Change last argument to use correct environment
# "prod": production environment
# "test": test environment
CMD ["/app/main", "-env", "prod"]
