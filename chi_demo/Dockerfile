ARG golang_version=1.19.3
ARG alpine_version=3.17
ARG golang_tag="$golang_version-alpine$alpine_version"

FROM golang:$golang_tag AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o build/server cmd/server/main.go

FROM alpine:$alpine_version AS cli

RUN apk add tar curl

ENV arch=amd64
ENV migrate_release=v4.15.2

RUN curl -LJO https://github.com/golang-migrate/migrate/releases/download/$migrate_release/migrate.linux-$arch.tar.gz

RUN tar xzf migrate.linux-$arch.tar.gz

FROM alpine:$alpine_version AS runner

WORKDIR /app

COPY --from=builder /app/build/server ./server

COPY --from=cli migrate ./migrate
