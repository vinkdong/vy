FROM golang:1.11.4-alpine3.8 as builder
WORKDIR /go/src/github.com/vinkdong/vy
ADD . .

# RUN \
#   apk add gcc build-base

RUN \
  go build .

FROM alpine:3.8

COPY --from=builder /go/src/github.com/vinkdong/vy/vy /usr/local/bin