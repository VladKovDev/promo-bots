FROM golang:1.24-alpine AS base

FROM base AS builder
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk update
RUN apk add --no-cache git build-base libc6-compat

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY apps/promobot/go.mod apps/promobot/go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application
COPY apps/promobot/ ./

# Install sqlc and generate code
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN sqlc generate

# Build the application
RUN CGO_ENABLED=0 go build -o /app/promobot-start /app/cmd/server/main.go

FROM base AS runner
WORKDIR /app

# Don't run production as root
RUN addgroup --system --gid 1001 promobotuser
RUN adduser --system --uid 1001 promobotuser

# Copy built application
COPY --from=builder --chown=promobotuser:promobotuser /app/promobot-start /usr/local/bin/promobot

# Copy configuration files
COPY --from=builder --chown=promobotuser:promobotuser /app/configs /app/configs

USER promobotuser

# Set APP_ENV to production by default (loads config.prod.yaml)
ENV APP_ENV=dev

EXPOSE 3000

CMD ["promobot"]