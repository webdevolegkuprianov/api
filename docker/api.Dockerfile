FROM golang:1.19-alpine as builder

WORKDIR /src
COPY . .

RUN go mod download

RUN go build ./cmd/api

CMD ["./api"]