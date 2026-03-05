# 构建阶段
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod  ./
RUN go mod tidy
COPY . .
RUN go build -o main .

# 运行阶段
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]