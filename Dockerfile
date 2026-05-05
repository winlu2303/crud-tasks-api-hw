FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o tasks-api cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/tasks-api .
EXPOSE 8080
CMD ["./tasks-api"]