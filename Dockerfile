FROM alpine:3.14.0

COPY bin/logger /bin/logger

CMD ["echo", "Container started!"]

ENTRYPOINT ["./bin/logger"]