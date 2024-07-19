FROM docker.io/golang:1.22-alpine AS build

RUN mkdir /src
RUN apk add git

ADD **/*.go /src
ADD ./go.* /src

WORKDIR /src

RUN go get -d -v -t
RUN GOOS=linux go build -v -o gobank
RUN chmod +x gobank

# ---
FROM docker.io/scratch AS production

COPY --from=build /src/gobank /usr/local/bin/gobank
EXPOSE 8080

CMD [ "gobank" ]