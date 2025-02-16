FROM golang:1.24.0 AS builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

RUN go mod tidy
RUN go mod download

RUN go build -o main .

FROM alpine:3.18

WORKDIR /root

RUN apk --no-cache add ca-certificates libc6-compat

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

RUN chmod +x /root/main

CMD ["./main"]
