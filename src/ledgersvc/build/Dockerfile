FROM alpine:3.21.0

RUN apk --no-cache add tzdata ca-certificates

COPY ./build/binaries/ /

USER 10001
