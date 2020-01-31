FROM golang:1.13.7

COPY . /var/www

WORKDIR /var/www

RUN go mod download

EXPOSE 3000

CMD ["go", "run", "/var/www/src/server.go"]
