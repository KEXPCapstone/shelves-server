FROM alpine
RUN apk add --no-cache ca-certificates
EXPOSE 443
COPY gateway /gateway
ENTRYPOINT ["/gateway"]