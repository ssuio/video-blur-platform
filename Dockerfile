# For Golang build
FROM golang:1.13.1-stretch AS builder
WORKDIR /server/
COPY ./server/ /server/

RUN GOOS=linux go build -o video-processing main.go

# For Server / App
FROM ubuntu:18.04
USER root
RUN apt-get update && \
apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
RUN apt-get update && \
apt-get install -y docker-ce docker-ce-cli containerd.io

COPY --from=builder /server/video-processing /video-processing-server

CMD ["/video-processing-server"]
