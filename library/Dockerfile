# first we build the go binary in a separate build stage
FROM golang:latest as builder
WORKDIR /go/src/github.com/KEXPCapstone/shelves-server/library
# install dependencies and build
COPY . .
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o librarysvc .

FROM scratch as runner
WORKDIR /root/
EXPOSE 80
COPY --from=builder /go/src/github.com/KEXPCapstone/shelves-server/library .
ENTRYPOINT ["./librarysvc"]