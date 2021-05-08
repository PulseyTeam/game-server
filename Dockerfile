FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY . .

RUN go get .

ENTRYPOINT go run main.go

EXPOSE 3000
