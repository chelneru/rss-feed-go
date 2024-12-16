FROM golang:1.23 AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o  main cmd/main.go

FROM alpine:latest AS run

COPY --from=build app app

EXPOSE 8080

WORKDIR /app


CMD ["./main"]