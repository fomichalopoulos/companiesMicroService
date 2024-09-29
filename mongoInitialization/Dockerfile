FROM golang:1.22.7-alpine3.20 AS build-img

RUN apk add --no-cache git

WORKDIR /tmp/Company

#We want to populate the module cache based on the go.{mod, sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download 

COPY . .

RUN go build -o build/mongoInit4Company .

FROM alpine:3.20
RUN apk add ca-certificates

COPY --from=build-img /tmp/Company/build/mongoInit4Company /app/mongoInit4Company
