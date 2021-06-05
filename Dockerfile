FROM golang:1.16-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go get .

RUN go build -o server .

FROM alpine:latest as deploy

WORKDIR /app/

COPY --from=builder /go/src/app/server .

EXPOSE 3000

ENTRYPOINT ["./server"]