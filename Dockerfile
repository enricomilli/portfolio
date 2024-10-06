FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install necessary build tools and dependencies
RUN apk add --no-cache git nodejs npm ca-certificates

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Install tailwindcss globally
RUN npm install -g tailwindcss

COPY go.mod go.sum ./

COPY . .

# Download Go dependencies
RUN go mod tidy


# Generate CSS
RUN tailwindcss -i site/assets/tailwind.css -o static/css/tailwind.css

# Generate templ files
RUN TEMPL_EXPERIMENT=rawgo templ generate

# Build the Go application
RUN go build -o ./bin/portfolio

# Start a new stage for a smaller final image
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the binary and necessary files from the builder stage
COPY --from=builder /app/bin/portfolio .
COPY --from=builder /app/static ./static

# Create a non-root user
RUN adduser -D appuser
USER appuser

CMD ["./portfolio"]