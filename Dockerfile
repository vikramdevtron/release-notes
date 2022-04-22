FROM golang:1.16.10-alpine3.13  AS build-env
RUN apk add --no-cache git gcc musl-dev
RUN apk add --update make
RUN go get github.com/google/wire/cmd/wire
WORKDIR /go/src/github.com/devtron-labs/release-notes
ADD . /go/src/github.com/devtron-labs/release-notes
RUN GOOS=linux make

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=build-env  /go/src/github.com/devtron-labs/release-info/release-notes .
CMD ["./release-notes"]