FROM docker.io/golang:1.22-alpine

WORKDIR /app

COPY ./go.* .
RUN go mod download

COPY ./**/*.go .

RUN go build -o ./gobank

EXPOSE 8080
CMD [ "./gobank" ]