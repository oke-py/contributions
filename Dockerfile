FROM golang:1.24 as builder

WORKDIR /go/src/github.com/oke-py/contributions
COPY . /go/src/github.com/oke-py/contributions

RUN make build

FROM gcr.io/distroless/base-debian12
COPY --from=builder /go/src/github.com/oke-py/contributions/bin/contribution /app
ENTRYPOINT [ "/app" ]
