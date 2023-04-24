FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/bin/order-management-service

FROM alpine:3.14

RUN apk update && apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/bin/order-management-service /app/

EXPOSE 8080

CMD ["./order-management-service"]