# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o gocrun .

# Final stage
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/gocrun /gocrun
EXPOSE 4000
CMD ["/gocrun"]
