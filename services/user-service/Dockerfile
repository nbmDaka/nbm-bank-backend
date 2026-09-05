FROM golang:1.26-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download


COPY . .


RUN go build -o user-service ./cmd/api



FROM alpine:latest


WORKDIR /app


COPY --from=builder /app/user-service .


COPY migrations ./migrations


EXPOSE 50051


CMD ["./user-service"]