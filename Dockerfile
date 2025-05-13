FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . ./
RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]