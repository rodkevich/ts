# build stage
FROM golang:1.17 AS builder
LABEL maintainer="golang linux builder"
ENV GOOS=linux

WORKDIR /go/src/github.com/rodkevich/ts
COPY . .
WORKDIR /go/src/github.com/rodkevich/ts/ticket/
RUN go mod download
RUN go build -ldflags="-s -w" ./cmd/ticket-service-server.go
CMD ["go", "test"]

# server stage
FROM scratch
LABEL maintainer="ivan.rodkevich/ticket-service-server"
COPY --from=builder /go/src/github.com/rodkevich/ts/ticket/ticket-service-server .
EXPOSE 8081
CMD ["./ticket-service-server"]
