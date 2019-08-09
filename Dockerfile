# STEP 1 build executable binary
FROM golang:1.12 as builder

COPY . $GOPATH/src/httpserver

#build the binary
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go install httpserver

# STEP 2 build a small image
# start from scratch

FROM alpine:latest
# Copy our static executable
COPY --from=builder /go/bin/httpserver /go/bin/httpserver

ENTRYPOINT ["/go/bin/httpserver"]
