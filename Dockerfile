FROM golang:1.14 AS build
WORKDIR /go/src/github.com/jdharmon/http-echo
COPY * ./
RUN go build

FROM scratch
COPY --from=build /go/src/github.com/jdharmon/http-echo /http-echo
ENTRYPOINT [ "/http-echo" ]