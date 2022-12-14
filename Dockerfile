FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["/app/main"]
