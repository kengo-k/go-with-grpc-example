FROM golang:1.19-alpine as builder
COPY volumes /root/src
WORKDIR /root/src
RUN go build cmd/server/main.go;

FROM alpine:3.16.2
WORKDIR /root
COPY --from=builder /root/src/main /root
CMD ["./main"]
