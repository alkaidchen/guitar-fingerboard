# builder
FROM golang:1.24 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o guitar-fingerboard .

# runner
FROM alpine

ENV TZ Asia/Shanghai

RUN apk add tzdata && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && apk del tzdata

WORKDIR /app

COPY --from=builder /app/guitar-fingerboard /usr/local/bin/guitar-fingerboard

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/guitar-fingerboard"]
