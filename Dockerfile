# FROM golang:1.13-alpine
# WORKDIR /app
# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY . .
# RUN go build -o /go/bin/main
# EXPOSE 9090
# ENTRYPOINT ["/go/bin/main"]

FROM golang:1.13-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /go/bin/main

FROM alpine:3.5
WORKDIR /app
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

COPY --from=builder go/bin/main .  
COPY ./config/config.yaml ./config/

#expose port
EXPOSE 9090
ENTRYPOINT ["./main"]

