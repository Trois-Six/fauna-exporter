FROM alpine

COPY fauna-exporter .

ENTRYPOINT ["/fauna-exporter"]
EXPOSE 80
