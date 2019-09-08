FROM golang:1.12.9-alpine

ENV GO111MODULE=on
ENV GOROOT=/usr/local/go
ENV GOPATH=/go/source

WORKDIR /go/source/simpleserver
COPY . .
RUN go mod download
RUN go build -o simpleserver .

ENTRYPOINT [ "./simpleserver" ]
