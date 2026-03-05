# 使用轻量级 Alpine 镜像
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码并编译
COPY . .
RUN go build -o main ./cmd/api/main.go

# 运行镜像
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
