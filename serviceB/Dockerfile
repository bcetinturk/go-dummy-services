FROM golang:1.20-buster

WORKDIR /app
COPY . .
RUN go build -o serviceB

EXPOSE 8081

CMD [ "./serviceB" ]
