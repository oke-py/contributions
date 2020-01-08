FROM golang:1.13 as builder

WORKDIR /go/src/github.com/oke-py/contributions
COPY . /go/src/github.com/oke-py/contributions

RUN make build

FROM gcr.io/distroless/base
COPY --from=builder /go/src/github.com/oke-py/contributions/bin/contribution /app
ENTRYPOINT [ "/app" ]
