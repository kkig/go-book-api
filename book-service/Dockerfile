# syntax=docker/dockerfile:1

FROM golang:1.20.3-bullseye AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o /book-service

# Run the tests in the container
# FROM build-stage AS run-test-stage

# RUN go test -v ./...

# Deploy app binary into lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /book-service /book-service

# To bind to TCP port, runtime param mus be supplied to the docer command.
# Optionally document in Dockerfile what ports app will listen.
EXPOSE ${HTTP_PORT}

USER nonroot:nonroot

CMD ["/book-service"]
