FROM golang:1.11.5-alpine3.8 AS builder

# ARG gRPCRole=client
# ARG gRPCRole=server
ARG gRPCRole=PleaseInputRole

# basic packages needed
RUN set -eux; \
        apk update && \
        apk add --no-cache --virtual .build-deps gcc libc-dev git

# env for go workdir and mod vendor
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/github.com/jimweng/pkgserver/
ENV GO111MODULE=on
# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR
# copy files into workdir
ADD . $GO_WORKDIR
RUN go mod vendor
RUN pwd
RUN cd ${gRPCRole};go build -o server -i main.go

FROM alpine:3.8
# copy binary into container
ARG gRPCRole=PleaseInputRole
COPY --from=builder /go/src/github.com/jimweng/pkgserver/${gRPCRole}/server .

CMD ["./server"]