FROM golang:alpine as builder

WORKDIR /app

# Copy Go module files
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy source files
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./build/api ./cmd/api/main.go

# FROM alpine:3.14.10
FROM scratch

EXPOSE 8080

COPY --from=builder /app/build/api .

ENV GOGC 1000
ENV GOMAXPROCS 3

ENTRYPOINT ["/api"]