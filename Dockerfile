FROM golang:1.16-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/prweb cmd/prweb/main.go
ENTRYPOINT [ "bin/prweb" ]
