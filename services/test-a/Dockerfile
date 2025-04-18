FROM golang:1.24-alpine

WORKDIR /app

COPY go.work ./
COPY ./services/test-a ./services/test-a
COPY ./libs ./libs

WORKDIR /app/services/test-a

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o test-service ./cmd/api

EXPOSE 8081

CMD ["./test-service"]
