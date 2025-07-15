# Build stage
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .

# Initialize module if needed
RUN test -f go.mod || go mod init example.com/ocean
RUN go mod tidy
RUN CGO_ENABLED=1 go build -o ocean ./cmd/ocean-demo

# Run stage
FROM debian:bullseye-slim
RUN apt-get update && \
    apt-get install -y libgl1-mesa-glx libx11-dev libxi6 \
                       libxcursor1 libxrandr2 libxinerama1 libxxf86vm1 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/ocean /usr/local/bin/ocean
CMD ["/usr/local/bin/ocean"]
