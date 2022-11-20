FROM golang:1.19.1-alpine3.16

RUN apk update && apk upgrade --no-cache
RUN apk add git --no-cache

WORKDIR /app

COPY . backend

WORKDIR /app/backend

RUN go mod tidy

RUN go build -o calendario-medico-api.bin  main.go

EXPOSE 8080

CMD ["./calendario-medico-api.bin"]