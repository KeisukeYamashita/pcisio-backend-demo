FROM golang:1.12 AS builder
LABEL MAINTAINER=KeisukeYamashita<19yamashita15@gmail.com>
ENV GO111MODULE on

WORKDIR /go/src/github.com/KeisukeYamashita/pcisio-server
COPY . .
RUN     go install -v server/server.go &&\
        make build

FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/KeisukeYamashita/pcisio-server/bin/server /go/bin/pcisio-server

EXPOSE 5050/tcp
CMD ["/go/bin/pcisio-server"]