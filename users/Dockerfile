FROM golang:alpine

WORKDIR /go/src/auth-service

COPY . .

RUN apk update && apk add git

# RUN echo $GOPATH

RUN go get -d -v ./...

#get dependancies
# RUN go get -d -v

# CMD ["go","run","main.go"]

EXPOSE 9000
