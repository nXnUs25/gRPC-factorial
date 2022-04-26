# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /factorial

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

ENV GRPC_PORT=5100

EXPOSE $GRPC_PORT

RUN go build -o calculater_server rpc_factorial/server/*.go 

CMD [ "./calculater_server" ]
