FROM docker.io/golang:1.22-alpine AS build

RUN mkdir /src
RUN apk add git

ADD ./*.go /src
ADD ./cmd /src/cmd/
ADD ./config /src/config/
ADD ./internals /src/internals/
ADD ./_common /src/_common/
ADD ./provider /src/provider/
ADD ./go.* /src

WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux go build -v -o go-api
RUN chmod +x go-api

# ---
FROM docker.io/golang:1.22-alpine AS production

COPY --from=build /src/go-api /usr/local/bin/go-api
EXPOSE 8080

CMD [ "go-api" ]