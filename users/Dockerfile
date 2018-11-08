FROM golang:alpine
WORKDIR /go/src/app
COPY . .
RUN apk update && apk add git

# RUN go get -d -v ./...

#get dependancies
RUN go get -d -v

CMD ["go","run","main.go"]

EXPOSE 9000
