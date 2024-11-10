FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN go build -o transhabit

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/ .

RUN chmod +x transhabit

EXPOSE 8000
CMD ["./transhabit"]
