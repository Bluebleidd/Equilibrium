
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o equilibrium .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/equilibrium .
COPY config.json .

EXPOSE 8000

CMD ["./equilibrium"]