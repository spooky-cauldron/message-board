# build step
FROM golang:1.21 AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /message-board-main


# output step
FROM alpine:3.18 AS main

WORKDIR /

COPY --from=build /message-board-main /message-board-main

ENV HOST=":5000"

EXPOSE 5000

ENTRYPOINT [ "/message-board-main" ] 
