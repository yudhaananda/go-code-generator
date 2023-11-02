FROM golang:1.18.0-alpine
RUN apk update && apk add -U --no-cache openssh git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go mod download
RUN go build -o main .
CMD ["/app/main"]