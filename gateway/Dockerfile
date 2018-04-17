# first we build the go binary in a separate build stage
FROM golang:latest as builder
WORKDIR /go/src/github.com/KEXPCapstone/shelves-server/gateway
# install dependencies and build
COPY . .
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myproject .
# RUN CGOENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o /gatewaysvc . && ls -la /gatewaysvc

FROM alpine:latest as runner
RUN apk add --no-cache ca-certificates
WORKDIR /root/
EXPOSE 443
COPY --from=builder /go/src/github.com/KEXPCapstone/shelves-server/gateway .
ENTRYPOINT ["./myproject"]
# COPY gatewaysvc /gatewaysvc
# RUN ls -l /gatewaysvc && echo "boop"
# ENTRYPOINT ["/gatewaysvc"]