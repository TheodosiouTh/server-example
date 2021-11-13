FROM golang:1.16-alpine

ADD . /go/src/server-example
WORKDIR /go/src/server-example

COPY go.mod ./
RUN go mod download

COPY * ./

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]