
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/api cmd/api/main.go


FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/api /app/api

EXPOSE 8080

CMD ["./api"]
