# ── Stage 1: Build ──
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git (needed for go mod download)
RUN apk add --no-cache git

# Copy go.mod and go.sum first for layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o travelsphere .

# ── Stage 2: Runtime ──
FROM alpine:3.18

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/travelsphere .

# Copy views, static assets, and config
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static
COPY --from=builder /app/conf ./conf

# Expose port
EXPOSE 8080

# Run the app
CMD ["./travelsphere"]
