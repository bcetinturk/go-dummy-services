FROM golang:1.20-buster

WORKDIR /app
COPY . .
RUN go build -o serviceA

EXPOSE 8080

CMD [ "./serviceA" ]
