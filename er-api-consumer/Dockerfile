FROM golang:1.18.1-buster as builder
WORKDIR /app
COPY . ./
RUN go build -o er-api-consumer .

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app /app
CMD ["/app/er-api-consumer"]
