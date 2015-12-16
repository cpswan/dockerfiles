FROM python:2

MAINTAINER @cpswan

# install utility packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    sudo curl wget vim

# install ansible from pip
RUN pip install ansible

# add an ansible user and allow it to sudo
RUN useradd -m ansible &&\
    echo 'ansible ALL=NOPASSWD: ALL' > /etc/sudoers.d/ansible

USER ansible

#default command
CMD /bin/bash

#working directory
WORKDIR /home/ansible

#example use: 'sudo docker run -it cpswan/ansible'
