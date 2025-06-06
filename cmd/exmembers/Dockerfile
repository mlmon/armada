# Build stage
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Run tests
RUN go test ./...

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o exmembers ./cmd/exmembers

# Final stage
FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/exmembers /exmembers

ENV PORT="7946"
ENV JOIN_ADDR=""

EXPOSE 7946

ENTRYPOINT ["/exmembers"]