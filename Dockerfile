# For Golang build

# For Server / App
FROM ubuntu:18.04
USER root
RUN echo "nameserver 8.8.8.8" | tee /etc/resolv.conf > /dev/null
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

