# build binary
FROM golang:1.10-alpine AS build

WORKDIR /go/src/srv

COPY . .

RUN pwd && ls -alrt

RUN go build srv

# copy to dest image
FROM golang:1.10-alpine

WORKDIR /opt/srv

RUN mkdir /hello_log

COPY --from=build /go/src/srv/srv /opt/srv/srv

CMD ["/opt/srv/srv"]
