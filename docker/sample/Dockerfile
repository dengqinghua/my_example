# docekr build -t dengqinghua/nginx .
# docker run -d -p 80 --name website -v $PWD/website:/var/www/html/website dengqinghua/nginx nginx
FROM ubuntu_dengqinghua
MAINTAINER dengqinghua "dengqinghua.42@gmail.com"
RUN apt-get -yqq update; apt-get -yqq install nginx
RUN mkdir -p /var/www/html/website
COPY global.conf /etc/nginx/conf.d/
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
