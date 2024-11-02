FROM golang:1.23.1-alpine AS builder-base
RUN apk add build-base


FROM builder-base AS build-stage
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /entrypoint

FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --chown=nonroot --from=build-stage /entrypoint /entrypoint
EXPOSE 5138
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
