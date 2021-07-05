FROM golang:1.16-alpine as builder

COPY . /build

WORKDIR /build

ENV GO116MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/logger main.go

# ENTRYPOINT ["./bin/logger"]

## Second Stage

From alpine:3.14.0

COPY --from=builder build/bin/logger /bin

CMD ["./bin/logger"]