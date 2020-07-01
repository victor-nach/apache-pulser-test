# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Iheanacho Victor nonso"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# RUN go mod tidy

# Download dependencies.
RUN go mod download

# Copy the source to the working directory in the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin .

# Start a new stage from scratch
FROM alpine:latest
#RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin .
#COPY --from=builder /app/.env .

#Command to run the executable
CMD ["./bin"]
