FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

