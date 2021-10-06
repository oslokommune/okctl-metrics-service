FROM golang:1.16 AS build

WORKDIR /go/src

ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go get -d -v ./...

COPY specification.yaml .
COPY main.go .

COPY ./pkg ./pkg

RUN go build -a -installsuffix cgo -o service .

FROM alpine:3.14 AS runtime

EXPOSE 3000/tcp
ENV GIN_MODE=release

RUN adduser -D appuser
USER appuser

COPY --from=build /go/src/service ./

ENTRYPOINT ["./service"]
