FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o transfer_system_app ./cmd/.

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/transfer_system_app .

EXPOSE 8080

CMD ["./transfer_system_app"]
