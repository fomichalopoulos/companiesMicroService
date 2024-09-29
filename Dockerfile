FROM golang:1.22.7-alpine3.20 AS build-img
#FROM golang:1.22.7-bullseye AS build-img

RUN apk add --no-cache git alpine-sdk musl-dev 

WORKDIR /tmp/Company

#We want to populate the module cache based on the go.{mod, sum} files.
COPY go.mod .
COPY go.sum .

#RUN apt-get update && apt-get install -y librdkafka-dev && \
RUN  go mod download 
#    go get github.com/confluentinc/confluent-kafka-go/kafka && \
#    go mod tidy

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o build/main -tags musl .

FROM alpine:3.20
RUN apk add ca-certificates 

COPY --from=build-img /tmp/Company/build/main /app/main
#COPY ./api/ui /app/api/ui

EXPOSE 8080

CMD ["/app/main"]

