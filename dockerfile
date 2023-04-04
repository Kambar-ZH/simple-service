FROM golang:1.20.1-alpine
RUN apk add git
RUN apk add --no-cache make
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN pwd
RUN go build -o main cmd/main.go
CMD ["/app/main"]