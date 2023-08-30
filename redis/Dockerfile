FROM debian:11-slim

RUN apt-get update
RUN apt-get install -y redis-server

EXPOSE 6379/tcp
WORKDIR /data

ENTRYPOINT ["redis-server"]
