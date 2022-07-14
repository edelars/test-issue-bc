FROM golang:1.18-alpine AS builder
WORKDIR /build
COPY . .
RUN go build   -ldflags '-s -w -extldflags "-static"' -o build/exe ./cmd/main.go

FROM alpine:latest
WORKDIR /build
COPY --from=builder /build /build

EXPOSE 12211

CMD ["./build/exe"]