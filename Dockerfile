# Base image, golang 1.18
FROM golang:1.19.3
WORKDIR /workspace
# Copy all files into the image
COPY . .
# Run go mod
RUN go mod download
# Expose ports
EXPOSE 8000
# Run Go program, just like locally
CMD ["go","run","cmd/playoff_bracket/main.go"]