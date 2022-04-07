FROM golang:1.17-alpine AS builder
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o dummy-online-shop .

FROM alpine:latest
WORKDIR /bin
COPY --from=builder /go/src/app/dummy-online-shop .
CMD dummy-online-shop
