# NOTE: this Dockerfile is configured for use with docker-compose
# and will only work properly if called from the context specified in
# docker-compose.yml
# first we build the go binary in a separate build stage
FROM golang:latest as builder
WORKDIR /go/src/github.com/KEXPCapstone/shelves-server/
# install dependencies and build
COPY . .
WORKDIR /go/src/github.com/KEXPCapstone/shelves-server/library
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o librarysvc .

FROM scratch as runner
WORKDIR /root/
EXPOSE 80
COPY --from=builder /go/src/github.com/KEXPCapstone/shelves-server/library .
ENTRYPOINT ["./librarysvc"]
