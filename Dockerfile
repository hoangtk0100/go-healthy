# Build stage
FROM golang:1.19.3-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY app.env.db .
COPY start.sh .
RUN chmod +x start.sh
COPY wait-for.sh .
RUN chmod +x wait-for.sh
COPY db/migration ./db/migration

EXPOSE 8088
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]