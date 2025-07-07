FROM golang:1.20 AS builder

WORKDIR /app

COPY src/ .

RUN go mod init go-hello-world
RUN go build -o hello-world .

FROM gcr.io/distroless/base

COPY --from=builder /app/hello-world /hello-world

CMD ["/hello-world"]