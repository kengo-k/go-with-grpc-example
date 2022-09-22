FROM golang:1.19-alpine

RUN \
  apk update; \
  apk add git; \
  apk add protoc; \
  apk add make; \
  apk add bash; \
  apk add jq;\
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 ;\
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1; \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7

WORKDIR /volumes

CMD ["/bin/ash"]