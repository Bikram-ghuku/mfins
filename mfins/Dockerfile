FROM golang:1.23.1 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o ./build --tags "fts5" -a -ldflags '-linkmode external -extldflags "-static"' .

FROM alpine:latest AS app
RUN apk --no-cache add tzdata ca-certificates bash
ENV TZ="Asia/Kolkata"
WORKDIR /app
COPY --from=builder /src/build .
CMD ["./build"]