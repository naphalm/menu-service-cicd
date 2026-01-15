# =========================
# Builder Stage
# =========================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# copy go mod files
COPY app/go.mod app/go.sum ./
RUN go mod download

# copy application source
COPY app/ .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o menu-service

# =========================
# Runtime Stage
# =========================
FROM alpine:3.19

WORKDIR /app

# install Postgres client and certificates
RUN apk add --no-cache postgresql-client ca-certificates bash

# copy binary
COPY --from=builder /app/menu-service .

# copy migration script
COPY scripts/migrate.sh ./scripts/migrate.sh
RUN chmod +x ./scripts/migrate.sh

# expose port
EXPOSE 3000

# run migrations and then the Go service
ENTRYPOINT ["sh", "-c", "./scripts/migrate.sh && ./menu-service"]
