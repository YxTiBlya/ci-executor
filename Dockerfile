FROM golang:1.21.0

WORKDIR /app

COPY cmd cmd/
COPY internal internal/
COPY transport transport/

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY config.yaml config.yaml

RUN go build cmd/main/main.go

STOPSIGNAL SIGTERM

CMD ["./main", "--cfg", "config.yaml"]