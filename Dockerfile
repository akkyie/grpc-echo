FROM golang:latest AS builder

WORKDIR /go/src/github.com/akkyie/grpc-echo
COPY . .
RUN go get ./...
RUN make CGO_ENABLED=0 GOOS=linux GOARCH=amd64 build

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/akkyie/grpc-echo/build/grpc-echo /grpc-echo
EXPOSE 50051
ENTRYPOINT ["/grpc-echo"]
