FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o auth-service ./main.go && ls -l auth-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/auth-service .

EXPOSE 8080

CMD ["./auth-service"]