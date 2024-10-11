FROM alpine:latest as build
RUN apk --no-cache add \
    lua5.3 lua5.3-dev luarocks5.3 \
    build-base zip
WORKDIR /src
COPY ./example /src
RUN luarocks-5.3 make --pack-binary-rock

FROM golang:latest as go-build
WORKDIR /src
COPY ./example-go/src /src
RUN go build -o .

FROM kong:3.7.1-ubuntu
USER root
RUN apt-get update && apt-get install -y unzip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
ENV KONG_DECLARATIVE_CONFIG=/etc/kong/kong.yaml
ENV KONG_DATABASE=off

COPY --from=build /src/kong-plugin-example-0.1.0-1.all.rock /tmp/
COPY --from=go-build /src/example-go /usr/local/bin/example-go
RUN luarocks install lua-resty-http && luarocks install luajson --dev
RUN luarocks install /tmp/kong-plugin-example-0.1.0-1.all.rock

ADD ./environments/kong.yaml /etc/kong/kong.yaml
ADD ./environments/kong.conf /etc/kong/kong.conf

USER kong