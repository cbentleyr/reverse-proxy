FROM golang:1.22

WORKDIR /app

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o /pong ./cmd

EXPOSE 8002

CMD ["/pong"]