# --- BUILD STAGE ---
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ai-hr-backend ./cmd/main.go

# --- RUN STAGE ---
FROM alpine:latest AS release

COPY --from=builder /app/ai-hr-backend /app/ai-hr-backend/
COPY config.yaml /etc/ai-hr-backend/

WORKDIR /app/ai-hr-backend/

EXPOSE 8080

CMD ["./ai-hr-backend"]