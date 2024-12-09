# Base image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy files
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 3000

# Run the application
CMD ["./main"]