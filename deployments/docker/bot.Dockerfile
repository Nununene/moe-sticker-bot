# Build stage
FROM golang:1.23 AS builder

WORKDIR /src

COPY . ./

RUN go build -o /moe-sticker-bot cmd/moe-sticker-bot/main.go

# Install the pure-Python emoji package outside the runtime image. The runtime
# image's system Python has a mismatched Expat library, which breaks pip.
FROM python:3.12-slim AS emoji-deps

RUN python -m pip install --no-cache-dir --target=/opt/msb-python-packages emoji==2.14.1

# Runtime stage
FROM ghcr.io/star-39/moe-sticker-bot:latest

COPY --from=emoji-deps /opt/msb-python-packages /opt/msb-python-packages
ENV PYTHONPATH=/opt/msb-python-packages

COPY --from=builder /moe-sticker-bot /moe-sticker-bot

ENTRYPOINT ["/moe-sticker-bot"]
