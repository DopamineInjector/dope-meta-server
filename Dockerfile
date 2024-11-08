FROM golang:1.23.1 AS builder-base
RUN apt update && apt install build-essential -y

FROM debian:bookworm AS release-base
RUN apt update && apt install sqlite3 -y

FROM builder-base AS builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o /entrypoint

FROM release-base AS release
WORKDIR /
COPY --from=builder /entrypoint /entrypoint
EXPOSE 5138
ENTRYPOINT ["/entrypoint"]
