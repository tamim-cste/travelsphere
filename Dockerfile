# ── Stage 1: Build ──
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o travelsphere .

# ── Stage 2: Runtime ──
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/travelsphere .
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static
COPY --from=builder /app/conf ./conf

EXPOSE 8080

CMD ["./travelsphere"]

