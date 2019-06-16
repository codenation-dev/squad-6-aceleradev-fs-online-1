FROM golang:1.12-alpine

RUN apk add --no-cache make git \
  && addgroup -g 1000 cn \
  && adduser -S -u 1000 cn cn \
  && mkdir -p /go/src/github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend \
  && chown -R cn:cn /go/src/github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend \
  && go get -u github.com/golang/dep/cmd/dep

USER cn
