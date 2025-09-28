# Build stage
FROM golang:1.25.1 AS builder
WORKDIR /app
COPY . .
RUN make build

# Deploy stage
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/server /app/server
COPY --from=builder /app/data /app/data
EXPOSE 8080
CMD ["/app/server"]
