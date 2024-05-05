FROM golang:1.22

WORKDIR /app

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o /ping ./cmd

EXPOSE 8001

CMD ["/ping"]