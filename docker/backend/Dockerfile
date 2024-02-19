FROM golang:1.21.7-bookworm as builder
WORKDIR /build

RUN apt-get update && apt-get install -y git

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
WORKDIR /build/cmd
RUN CGO_ENABLED=0 go build -o app .

# Create final image
FROM alpine as runtime
WORKDIR /app
COPY --from=builder /build/cmd/app .

EXPOSE 3000
ENTRYPOINT ["/app/app"]