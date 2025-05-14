# ============================
# Stage 1: Build the Go binary
# ============================
FROM golang:1.23.6 AS builder

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and .env file
COPY . .
COPY .env /app/.env

# Build the webserver binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/webserver ./cmd/webserver

# ============================
# Stage 2: Run the application
# ============================
FROM alpine:3.18

WORKDIR /app

# Copy the compiled binary and .env
COPY --from=builder /app/bin/webserver .
COPY --from=builder /app/.env .env

# Copy the templates and static assets
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/monthlyreadings ./monthlyreadings
COPY --from=builder /app/cmd/MadameAI ./MadameAI

# Expose port
EXPOSE 8080

# Run the application
CMD ["./webserver"]
