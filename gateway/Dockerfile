# first we build the go binary in a separate build stage
FROM golang:latest as builder
WORKDIR /go/src/github.com/KEXPCapstone/shelves-server/gateway
# install dependencies and build
COPY . .
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gatewaysvc .

FROM alpine:latest as runner
RUN apk add --no-cache ca-certificates
WORKDIR /root/
EXPOSE 443 80
COPY --from=builder /go/src/github.com/KEXPCapstone/shelves-server/gateway .
ENTRYPOINT ["./gatewaysvc"]
