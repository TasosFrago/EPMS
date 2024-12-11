FROM golang:1.23-bookworm AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-api

EXPOSE 8080

CMD ["/build/go-api"]
