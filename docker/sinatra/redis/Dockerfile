# docker build -t dengqinghua/redis .
# docker run -d -p 6379 --name redis dengqinghua/redis
FROM ubuntu_dengqinghua
MAINTAINER dengqinghua "dengqinghua.42@gmail.com"
RUN apt-get update -yqq; apt-get install -yqq redis-server redis-tools
EXPOSE 6379
ENTRYPOINT ["/usr/bin/redis-server"]
