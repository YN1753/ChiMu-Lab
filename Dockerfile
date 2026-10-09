# ==========================================
# Stage 1: Build Frontend (Vue 3 + Vite)
# ==========================================
FROM node:22-alpine AS frontend-builder
WORKDIR /build/web

COPY web/package*.json ./
RUN npm ci

COPY web/ ./
RUN npm run build

# ==========================================
# Stage 2: Build Backend (Go 1.24+)
# ==========================================
FROM golang:1.24-alpine AS backend-builder
WORKDIR /build

RUN apk add --no-cache git ca-certificates tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 编译纯 Go CGO-Free 静态二进制
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o chimu-server main.go

# ==========================================
# Stage 3: Minimal Production Image
# ==========================================
FROM alpine:3.21
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 从构建阶段复制二进制与前端构建产物
COPY --from=backend-builder /build/chimu-server /app/chimu-server
COPY --from=frontend-builder /build/web/dist /app/web/dist

# 创建数据存储目录
RUN mkdir -p /app/data

EXPOSE 8080

VOLUME ["/app/data"]

ENV GIN_MODE=release \
    PORT=8080 \
    DB_PATH=/app/data/chimu.db

CMD ["/app/chimu-server"]
