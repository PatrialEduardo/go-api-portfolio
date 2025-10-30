FROM golang:1.24 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/main.go

# imagem final
FROM debian:stable-slim
WORKDIR /app

# instalar psql para rodar migration
RUN apt-get update && \
    apt-get install -y --no-install-recommends postgresql-client ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /app/server /app/server
COPY db/init.sql /app/db/init.sql

ENV PORT=8080
EXPOSE 8080

CMD ["/app/server"]
