# Docker.Redis.Server
FROM centos
MAINTAINER yqjdcyy <yqjdcyy@gmail.com>
COPY . /
RUN chmod +x server
CMD ./server -port=3000 -addr=172.17.0.1:8030
