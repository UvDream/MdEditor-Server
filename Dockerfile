FROM golang:1.18.4-buster AS builder
WORKDIR /go/src/app
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o server main.go

FROM alpine:3.15.4
WORKDIR /opt
COPY --from=builder /go/src/app/server /opt
COPY config.production.yaml /opt
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
 chmod +x /opt/server
ENTRYPOINT ["/opt/server", "-config=/opt/config.production.yaml"]