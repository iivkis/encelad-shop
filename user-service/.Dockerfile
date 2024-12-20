FROM golang:alpine3.21

WORKDIR /

COPY ./shared/go.mod ./shared/go.sum /shared/

RUN cd /shared && go mod download

COPY ./shared /shared

COPY ./user-service/go.mod ./user-service/go.sum /service/

RUN cd /service && go mod download

COPY ./user-service /service

RUN cd /service && go build -o app ./cmd

CMD ["./service/app"]
