FROM node:0.10-slim

MAINTAINER @cpswan

# install utility packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    sudo curl wget vim

# install Azure CLI
RUN npm install azure-cli -g

# add an azurecli user and allow it to sudo
RUN useradd -m azurecli &&\
    echo 'azurecli ALL=NOPASSWD: ALL' > /etc/sudoers.d/azurecli

USER azurecli

#default command
CMD /bin/bash
