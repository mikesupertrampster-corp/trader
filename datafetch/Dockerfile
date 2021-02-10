FROM golang:1.15 as builder
WORKDIR /workspace
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:3.13
RUN apk add --no-cache ca-certificates
COPY --from=builder /workspace/datafetch /

ENTRYPOINT ["/datafetch"]