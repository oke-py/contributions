FROM golang:1.13 as builder

WORKDIR /go/src/github.com/oke-py/contribution
COPY . /go/src/github.com/oke-py/contribution

RUN make build

FROM gcr.io/distroless/base
COPY --from=builder /go/src/github.com/oke-py/contribution/bin/contribution /app
ENTRYPOINT [ "/app" ]
