FROM node:20

MAINTAINER cpswan

# install node-red
RUN npm install -g --unsafe-perm node-red

# expose port
EXPOSE 1880

# Set the default command to execute
# when creating a new container
CMD ["/usr/local/bin/node-red"]
