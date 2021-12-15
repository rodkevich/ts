# build stage
FROM golang:1.17 AS builder
LABEL maintainer="golang applictions linux builder"
#ENV GO111MODULE=auto
#ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /source
COPY . .
WORKDIR /source/ticket
RUN go mod download

RUN go run ./cmd/main.go
RUN go build -ldflags="-s -w" ./ticket/cmd/main.go

## server stage
#FROM scratch
#LABEL maintainer="ivan.rodkevich/app"
#
#COPY --from=builder /go/src/github.com/rodkevich/tbd/cmd/task5/server .
#
#EXPOSE 12300
#
#CMD ["./server"]