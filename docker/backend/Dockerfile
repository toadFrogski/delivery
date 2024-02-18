FROM golang:1.21.7-bookworm as builder
WORKDIR /build

RUN apt-get update && apt-get install -y git

# Build
COPY . .
WORKDIR /build/cmd
RUN CGO_ENABLED=0 go build -o app .

# Create final image
FROM alpine
WORKDIR /app
COPY --from=builder /build/cmd/app .

EXPOSE 3000
ENTRYPOINT ["/app/app"]