FROM golang:1.14

ENV GOPATH /usr/src/app/go
ENV TOKEN_FROM_PROMPT 1

ARG dir=$GOPATH/src/github.com/mrauer

WORKDIR $GOPATH/src/github.com/mrauer/freetube

RUN apt-get update && apt-get install -y ffmpeg

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN curl -Lo youtubedr.tar.gz https://github.com/kkdai/youtube/releases/download/v2.7.2/youtubedr_2.7.2_linux_amd64.tar.gz && tar -xvf youtubedr.tar.gz && chmod +x youtubedr && mv youtubedr /usr/local/bin

COPY . .
