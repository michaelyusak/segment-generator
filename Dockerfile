# First stage: build the Go binary
FROM golang:1.25.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o segment-generator .

FROM alpine:latest

COPY --from=builder /app/segment-generator /segment-generator

RUN chmod +x /segment-generator

EXPOSE 8080

CMD ["/segment-generator"]
