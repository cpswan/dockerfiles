FROM ubuntu:12.04

MAINTAINER cpswan
# Add universe repository to /etc/apt/sources.list
# we need it for nginx
RUN sed -i s/main/'main universe'/ /etc/apt/sources.list
# Update repos so that changes can take effect
RUN apt-get update
# Install nginx
RUN apt-get install -y nginx
# Turn off daemon mode
RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf
# Meetup customisation
RUN sed -i s/nginx/'Chicago Docker Meetup'/ /usr/share/nginx/www/index.html
# Expose web server
EXPOSE 80
# Run nginx
CMD /usr/sbin/nginx
