FROM golang:1.13.7

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 3000

ENTRYPOINT CompileDaemon --command="./main"