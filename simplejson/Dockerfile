FROM golang:1.15 as builder

WORKDIR /workspace
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.13
COPY --from=builder /workspace/ /

ENTRYPOINT ["/main"]