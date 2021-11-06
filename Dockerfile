FROM golang:1.17.2

USER root
WORKDIR /app

COPY . .

RUN go build -o /go-binary

EXPOSE 5001

CMD [ "/go-binary" ]

