# -------- Stage 1: Dev & Build --------
FROM golang:1.24-alpine as dev

# Install Air for dev hot reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.work ./
COPY ./services/test-a ./services/test-a
COPY ./libs ./libs

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o test-service ./services/test-a/cmd/api

EXPOSE 8081

# -------- Stage 2: Production --------
FROM alpine:latest AS prod

WORKDIR /app
COPY --from=dev /app/test-service .

CMD ["./test-service"]
