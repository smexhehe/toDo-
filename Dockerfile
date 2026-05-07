FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o todo-api ./cmd/todo

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/todo-api .

RUN mkdir -p data

EXPOSE 8080

CMD ["./todo-api"]
