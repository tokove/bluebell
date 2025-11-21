FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o bluebell_app .

FROM alpine:3.18

RUN apk add --no-cache bash netcat-openbsd

COPY wait-for.sh /
COPY templates /templates
COPY static /static
COPY config /config
COPY --from=builder /build/bluebell_app /

RUN chmod +x /wait-for.sh

EXPOSE 8081

ENTRYPOINT ["/wait-for.sh", "mysql8019:3306", "redis507:6379", "--", "/bluebell_app", "/config/config.yaml"]