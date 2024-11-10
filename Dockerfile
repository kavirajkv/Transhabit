FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /transhabit


FROM alpine:latest

WORKDIR /
COPY --from=builder /transhabit /transhabit

EXPOSE 8000

ENTRYPOINT ["/transhabit"]

