FROM golang:1.16-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /build

COPY . .

RUN go get .

RUN go build -o server .

FROM alpine:latest as alpine

WORKDIR /app/

COPY ./config ./config

COPY --from=builder /build/server .

EXPOSE 3000

ENTRYPOINT ["./server"]