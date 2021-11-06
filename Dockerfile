FROM golang:1.17.2

WORKDIR /app

COPY . .

RUN go build -o /go-bin

EXPOSE 5001

CMD [ "/go-bin" ]

