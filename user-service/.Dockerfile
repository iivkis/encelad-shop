FROM golang:alpine3.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o service -ldflags="-s -w" ./cmd

CMD ["./service"]