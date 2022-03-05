FROM golang:1.17 as builder

WORKDIR /go/src/github.com/oke-py/contributions
COPY . /go/src/github.com/oke-py/contributions

RUN make build

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/src/github.com/oke-py/contributions/bin/contribution /app
ENTRYPOINT [ "/app" ]
