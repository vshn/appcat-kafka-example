FROM docker.io/library/alpine:3.16 as runtime

RUN \
  apk add --update --no-cache \
    bash \
    curl \
    ca-certificates \
    tzdata

COPY appcat-kafka-example /usr/bin/appcat-kafka-example
ENTRYPOINT ["appcat-kafka-example"]

USER 65536:0
