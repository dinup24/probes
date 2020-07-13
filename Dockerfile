FROM golang:1.14-buster AS build
ENV GOPROXY=https://proxy.golang.org
WORKDIR /go/src/probes
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/probes .

FROM ubuntu:bionic
RUN apt-get update
RUN apt-get dist-upgrade -y
RUN mkdir /app
COPY --from=build /go/bin/probes /app/
WORKDIR /app
USER root
CMD ["./probes"]
