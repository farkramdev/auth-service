FROM golang:alpine
RUN mkdir -p $GOPATH/src/auth-service/users
WORKDIR $GOPATH/src/auth-service/users
COPY . $GOPATH/src/auth-service/users
RUN apk update && apk add git
RUN go get -d -v ./...
CMD ["go","run","main.go"]
EXPOSE 9000
