FROM golang:1.19.2-alpine

RUN apk update && apk --no-cache add git build-base

WORKDIR /go/src/github.com/takeruun/generate-pdf-go

COPY . .

RUN go install golang.org/x/tools/gopls@latest