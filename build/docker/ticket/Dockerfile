# build stage
FROM golang:1.17 AS builder
LABEL type="golang linux builder"
ENV GOOS=linux

WORKDIR /go/src/github.com/rodkevich/ts
COPY ticket/go.mod .
COPY ticket/go.sum .
RUN go mod download
COPY . .
WORKDIR /go/src/github.com/rodkevich/ts/ticket
RUN go build -ldflags="-s -w" ./cmd/ticket.go
CMD ["go", "test"]

# server stage
#FROM scratch
FROM golang:1.17
LABEL type="ivan.rodkevich/ticket-service-server"
COPY --from=builder /go/src/github.com/rodkevich/ts/ticket/config/ config/
COPY --from=builder /go/src/github.com/rodkevich/ts/ticket/ticket .

ENV CONFIG=docker

EXPOSE 8081
CMD ["./ticket"]
