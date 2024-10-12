# Stage 1: Build stage
FROM ubuntu:22.04 AS builder

# Install deno
RUN apt-get update && apt-get install -y curl unzip && \
    curl -fsSL https://deno.land/x/install/install.sh | sh && \
    mv /root/.deno/bin/deno /usr/local/bin/deno

# Set the working directory
WORKDIR /app

# Copy the hello.ts file into the container
COPY hello.ts .

# Compile hello.ts
RUN deno compile --allow-net hello.ts

# Stage 2: Final stage using distroless
FROM gcr.io/distroless/cc-debian12

# Copy the compiled binary from the builder stage
COPY --from=builder /app/hello /app/hello

# Set the working directory
WORKDIR /app

# Run the compiled binary
CMD ["./hello"]
