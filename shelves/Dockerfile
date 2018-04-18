# FROM scratch
# COPY shelves /shelves
# EXPOSE 80
# ENTRYPOINT ["/shelves"]

# first we build the go binary in a separate build stage
FROM golang:latest as builder
WORKDIR /go/src/github.com/KEXPCapstone/shelves-server/shelves
# install dependencies and build
COPY . .
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shelvessvc .

FROM scratch as runner
WORKDIR /root/
EXPOSE 80
COPY --from=builder /go/src/github.com/KEXPCapstone/shelves-server/shelves .
ENTRYPOINT ["./shelvessvc"]