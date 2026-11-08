# Build stage
FROM golang:1.23 AS builder

WORKDIR /src

COPY . ./

RUN go build -o /moe-sticker-bot cmd/moe-sticker-bot/main.go


# Runtime stage
FROM ghcr.io/star-39/moe-sticker-bot:latest

RUN apt-get update \
    && apt-get install -y --no-install-recommends python3-emoji \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /moe-sticker-bot /moe-sticker-bot

ENTRYPOINT ["/moe-sticker-bot"]
