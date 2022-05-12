# Start from golang base image
FROM golang:alpine as builder
LABEL maintainer="Sérgio Alves <juniorspse@gmail.com>"
# Install git.
RUN apk update && apk add --no-cache git
# Set the current working directory inside the container 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

EXPOSE 8080
CMD ["./main"]
