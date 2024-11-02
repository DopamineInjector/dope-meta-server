FROM golang:1.23.1-alpine as builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --chown=nonroot --from=builder /entrypoint /entrypoint
EXPOSE 5138
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
