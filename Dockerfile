FROM golang:1.19-alpine as builder
RUN apk update && apk add openssh-client gcc g++ musl-dev
WORKDIR /app
COPY . ./
RUN --mount=type=cache,target=/root/go/pkg/mod go mod download
RUN --mount=type=cache,target=/root/.cache/go-build go build -o search cmd/*.go

FROM alpine:latest  
RUN apk add ca-certificates supervisor
WORKDIR /app
COPY supervisord.conf /etc/supervisord.conf
COPY --from=builder /app/. ./
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
