FROM debian:11-slim

RUN apt-get update && apt-get upgrade
RUN apt-get install nginx -y


COPY nginx.conf /etc/nginx/nginx.conf
CMD ["nginx", "-g", "daemon off;"]

