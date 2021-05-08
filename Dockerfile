FROM golang:1.13

ENV GOPATH /usr/src/app/go
ARG dir=$GOPATH/src/github.com/mrauer
WORKDIR ${dir}

WORKDIR $GOPATH/src/github.com/mrauer/freetube

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
