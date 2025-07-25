# Build stage
# Use the same Debian release for build and runtime stages to avoid
# glibc version mismatches between the compiled binary and the base
# system. Bookworm ships with glibc >= 2.36 which matches the
# libraries used by the Go builder image.
FROM golang:1.20-bookworm AS builder
WORKDIR /app
COPY . .

# Install build dependencies for OpenGL bindings
RUN apt-get update && \
    apt-get install -y pkg-config libgl1-mesa-dev libx11-dev \
                       libxi-dev libxcursor-dev libxrandr-dev \
                       libxinerama-dev libxxf86vm-dev && \
    rm -rf /var/lib/apt/lists/*

# Initialize module if needed
RUN test -f go.mod || go mod init example.com/ocean
RUN go mod tidy
RUN CGO_ENABLED=1 go build -o ocean ./cmd/ocean-demo

# Run stage
FROM debian:bookworm-slim
RUN apt-get update && \
    apt-get install -y libgl1-mesa-glx libx11-dev libxi6 \
                       libxcursor1 libxrandr2 libxinerama1 libxxf86vm1 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/ocean /usr/local/bin/ocean
CMD ["/usr/local/bin/ocean"]
