FROM golang:1.13 AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 go build -mod=readonly -o /server

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /server /server

ENTRYPOINT /server

