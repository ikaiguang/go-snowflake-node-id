ARG WORK_DIR=/www/app
ARG GO_PATH=/www/go
ARG GO_PROXY=https://goproxy.cn,direct

FROM golang:1.18 as builder

ARG WORK_DIR
ARG GO_PATH
ARG GO_PROXY

WORKDIR $WORK_DIR

ENV GOPATH=$GO_PATH
ENV GOPROXY=$GO_PROXY
ENV CGO_ENABLED=0

COPY . .

#RUN ls -al
RUN go build -o ./bin/main ./cmd/main/snowflake-node-id.go

FROM alpine:latest as release

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add -U tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' > /etc/timezone

ARG WORK_DIR

WORKDIR $WORK_DIR

COPY --from=builder $WORK_DIR/bin/main $WORK_DIR/main

ENTRYPOINT ["./main"]

EXPOSE 8081 9091
