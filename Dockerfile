# Sử dụng base image cho Go
FROM golang:1.23-alpine AS builder

# Thiết lập thư mục làm việc
WORKDIR /app

# Copy file module và download các dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source code vào container
COPY . .

# Biên dịch ứng dụng
RUN go build -o main ./src/cmd/main.go

# Sử dụng image chính thức của PostgreSQL
FROM postgres:alpine AS postgres_base

# Sử dụng base image tối giản cho production
FROM alpine:latest

# Cài đặt các công cụ cần thiết
RUN apk add --no-cache bash curl postgresql-client

# Copy ứng dụng từ builder
COPY --from=builder /app/main /usr/local/bin/main

# Copy database từ postgres_base
COPY --from=postgres_base /usr/local/bin/docker-entrypoint.sh /usr/local/bin/
COPY --from=postgres_base /usr/local/bin/initdb /usr/local/bin/

# Copy migration files và cấu hình database
COPY ./src/infrastructure/migrations /migrations
COPY ./docker/docker-compose.yml /docker-compose.yml

# Thiết lập biến môi trường cho ứng dụng
ENV POSTGRES_USER=postgres \
    POSTGRES_PASSWORD=admin \
    POSTGRES_DB=relibcadb

# Expose các cổng cần thiết
EXPOSE 5432 8080

# Command để chạy ứng dụng
CMD ["bash", "-c", "/usr/local/bin/main & postgres -D /var/lib/postgresql/data"]