# Use the official Golang image as the base image
FROM golang:1.24-alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server/main.go

# Final stage
FROM alpine:3.22
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=go-builder /app/main .
# COPY --from=go-builder /app/internal/db/migration /app/internal/db/migration
COPY --from=go-builder /app/static /app/static
EXPOSE 8088
CMD ["./main"]