# Stage 1: Modules
FROM golang:1.22.4 as modules
WORKDIR /modules
COPY go.mod go.sum ./
RUN go mod download

# Stage 2: Builder
FROM golang:1.22.4 as builder
WORKDIR /app
COPY . .
COPY --from=modules /go/pkg /go/pkg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main server.go

# Stage 3: Final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
