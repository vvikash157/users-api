FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache libc6-compat
COPY --from=builder /app/main .
RUN chmod +x main
EXPOSE 8080
COPY .env .env
CMD ["./main"]


# Smaller Image: The final image is lightweight (~5MB from Alpine + binary).
# Faster Builds: go.mod is copied first to leverage Docker's caching mechanism.
# Security: The final image doesn’t include unnecessary files like source code.
# Portability: The binary is compiled statically (CGO_ENABLED=0), making it easy to run anywhere.  */
