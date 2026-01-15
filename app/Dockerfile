FROM golang:1.12-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o menu-service

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/menu-service .
COPY ../scripts/migrate.sh /app/migrate.sh
RUN chmox +x /app/migrate.sh

EXPOSE 3001
ENTRYPOINT [ "/app/menu-service" ]