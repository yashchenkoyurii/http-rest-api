FROM golang:latest as builder
COPY . /go/src/application
WORKDIR /go/src/application
ENV GO111MODULE="on"
RUN go mod vendor
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -o /go/bin/service ./cmd/apiserver/

FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service
CMD /go/bin/service
