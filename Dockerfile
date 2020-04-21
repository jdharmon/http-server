FROM golang:1.14 AS build
WORKDIR /app
COPY * ./
RUN CGO_ENABLED=0 go build -o http-server

FROM scratch
COPY --from=build /app/http-server /http-server
ENTRYPOINT [ "/http-server" ]