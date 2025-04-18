FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
