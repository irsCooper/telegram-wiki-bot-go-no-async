FROM golang:1.21.1-alpine3.18 AS builder

COPY  . /tg-bot-wikipedia
WORKDIR /tg-bot-wikipedia

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=0 /tg-bot-wikipedia/bin/bot .
COPY --from=0 /tg-bot-wikipedia/configs configs/


EXPOSE 80

CMD ["./bot"]
