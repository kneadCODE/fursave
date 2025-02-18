FROM golang:1.23-alpine AS builder

COPY . .

RUN for CMD in `ls cmd`; do \
    (CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o build/binaries/$CMD ./cmd/$CMD) \
    done

FROM alpine:3.21.3

RUN apk --no-cache add tzdata ca-certificates

COPY --from=builder ./go/build/binaries/ /

USER 1000
