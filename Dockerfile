FROM alpine

RUN apk --no-cache --no-progress add git ca-certificates tzdata make \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

COPY fauna-exporter .

ENTRYPOINT ["/fauna-exporter"]
EXPOSE 80
