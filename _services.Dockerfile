# Dockerfile References: https://docs.docker.com/engine/reference/builder/
############################
# STEP 1 build executable binary
############################
FROM golang:1.13.1-alpine AS builder

# Add svc ARGS
ARG SVC

# Create appuser.
RUN adduser -D -g '' appuser

# Copy project to working directory
RUN mkdir /app
COPY . /app/
WORKDIR /app/services/${SVC}/server

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main server.go plugin.go


############################
# STEP 2 build a small image
############################
FROM scratch

ARG SVC

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable.
COPY --from=builder /app/services/${SVC}/server/main /app/main
# Copy default config.
COPY --from=builder /app/core/config/.app.config.json.example /app/core/config/.app.config.json
# Copy local timezone.
COPY --from=builder /app/core/TZ-Jakarta /etc/localtime

WORKDIR /app

# Use an unprivileged user.
USER appuser

# Set env for registry
ENV MICRO_REGISTRY 'kubernetes'

# Run the binary.
ENTRYPOINT ["/app/main"]
