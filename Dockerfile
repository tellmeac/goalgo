FROM golang:1.21.5-alpine3.19
LABEL authors="Alexander Lipatov tellmeac@gmail.com"

WORKDIR ./app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./server ./cmd/server
RUN go build -o ./bot ./cmd/bot-publisher

CMD ./server & ./bot