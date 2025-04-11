# Step 1 - Build
FROM golang:alpine AS stage1

# Define the working directory
WORKDIR /app

# Copy the files go.mod and go.sum to the root of the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . . 

# Build the application
RUN go build -o /app/main .

# Step 2 - Final image optimized
FROM alpine:latest

# Define the working directory
WORKDIR /root/

# Copy the binary from the previous stage
COPY --from=stage1 /app/main .

# Copy config folder to the root of the working directory
COPY --from=stage1 /app/config ./root/config

# If you have to run in develop mode, copy the env file
COPY --from=stage1 /app/.env.development ./

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main", "-e", "development"]