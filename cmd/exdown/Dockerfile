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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o exdown ./cmd/exdown

# Final stage
FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/exdown /exdown

ENV SRC_URL=""
ENV DEST_PATH="/tmp/downloaded_file"

ENTRYPOINT ["/exdown"]