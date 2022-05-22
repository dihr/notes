# build stage
FROM golang:1.18-alpine as build

ADD . /go/src/github.com/dihr/app
WORKDIR /go/src/github.com/dihr/app
RUN go mod download \
    && go mod tidy
RUN CGO_ENABLED=0 go build -a -installsuffix main.go -o main

# final stage
FROM alpine
WORKDIR /app
COPY --from=build /go/src/github.com/dihr/app/main /app/main
ENTRYPOINT ./main