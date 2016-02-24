FROM python:2-slim

MAINTAINER @cpswan

# install utility packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    sudo curl wget vim

# install awscli and jpterm
RUN pip install awscli jmespath-terminal

# add an awscli user and allow it to sudo
RUN useradd -m awscli &&\
    echo 'awscli ALL=NOPASSWD: ALL' > /etc/sudoers.d/awscli

# enable command completion
RUN echo "complete -C '/usr/local/bin/aws_completer' aws" >> \
    /home/awscli/.bashrc

USER awscli

#default command
CMD /bin/bash

#example use: 'sudo docker run -it cpswan/awscli'
