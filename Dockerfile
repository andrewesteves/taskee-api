FROM golang:1.14 AS build

WORKDIR /go/src/taskee-api

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /main

COPY ./env.yaml ./env.yaml
COPY --from=build /go/src/taskee-api/main .

EXPOSE 80

CMD ["./main"]