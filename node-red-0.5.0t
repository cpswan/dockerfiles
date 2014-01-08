FROM ubuntu:12.04

MAINTAINER cpswan

# add universe repository to /etc/apt/sources.list
# we need it later as rlwrap is required by node.js
RUN sed -i s/main/'main universe'/ /etc/apt/sources.list

# make sure everything is up to date - update and upgrade
RUN apt-get update && apt-get upgrade -y

# install dependencies
RUN apt-get install python-software-properties python g++ make software-properties-common wget unzip -y

# add node.js repo
RUN add-apt-repository ppa:chris-lea/node.js -y

# another update so added repo can be used
RUN apt-get update

# install node.js
RUN apt-get install nodejs -y

# download node-red
RUN wget https://github.com/node-red/node-red/archive/0.5.0.zip

# unzip node-red into /opt
RUN unzip 0.5.0.zip -d /opt

# install node-red
RUN cd /opt/node-red-0.5.0 && npm install --production

# remove node-red archive
RUN rm 0.5.0.zip
