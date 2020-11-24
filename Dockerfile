FROM golang:1.15.5-buster AS builder
WORKDIR /go/src/github.com/jar-b/dsemver
COPY main.go go.mod go.sum ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /dsemver .

FROM scratch
COPY --from=builder /dsemver .
ENTRYPOINT ["/dsemver"]