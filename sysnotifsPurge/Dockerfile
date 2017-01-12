FROM golang:latest

ADD . /go/src/account-api-user
WORKDIR /go/src/account-api-user
RUN go get 

RUN go install 

ENTRYPOINT /go/bin/account-api-user

EXPOSE 5000
