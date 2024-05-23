FROM golang:1.22.0-alpine3.19 as builder 

WORKDIR /app

COPY  go.mod ./
COPY  go.sum ./
COPY internals ./internals
COPY controllers ./controllers
COPY bin ./bin
COPY views ./views
COPY emails ./emails
COPY main.go ./main.go

RUN go build -o go-scaffold


FROM alpine:3.19.1

WORKDIR /
COPY --from=builder /app/go-scaffold /go-scaffold
COPY .env /.env
COPY emails /emails
COPY static /static
COPY views /views

EXPOSE 8080

CMD [ "/go-scaffold" ]
