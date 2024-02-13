# STAGE 1: Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o main ./main.go

# STAGE 2: Run Stage
FROM alpine:3.19.0
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY db/migration ./db/migration

EXPOSE 8910
CMD ["nohup", "./main"]