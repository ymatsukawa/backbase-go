FROM golang:1.24.3-alpine3.22

ENV GOARCH=$GOARCH

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# tools for development
## https://github.com/air-verse/air?tab=readme-ov-file#installation-and-usage-for-docker-users-who-dont-want-to-use-air-image
RUN go install github.com/air-verse/air@v1.62.0
RUN go install github.com/go-delve/delve/cmd/dlv@v1.25.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0
