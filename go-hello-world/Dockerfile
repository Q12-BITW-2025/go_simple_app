FROM golang:1.20 AS builder

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src/ .

RUN go build -o hello-world .

FROM gcr.io/distroless/base

COPY --from=builder /app/hello-world /hello-world

CMD ["/hello-world"]