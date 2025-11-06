# Build stage
FROM golang:1.25.1-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Deploy stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/db/data.json /app/db/data.json

EXPOSE 8080
CMD [ "/app/main" ]
