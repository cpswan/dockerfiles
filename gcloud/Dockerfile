FROM python:2.7.9-slim

MAINTAINER @cpswan

# install utility packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    sudo curl wget vim openssh-client

# add an awscli user and allow it to sudo
RUN useradd -m gcloud &&\
    echo 'gcloud ALL=NOPASSWD: ALL' > /etc/sudoers.d/gcloud

# switch to gcloud user
USER gcloud

# install gcloud
RUN curl -o /tmp/google-cloud-sdk.tar.gz \
    https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz \
    && tar -C /home/gcloud -zxf /tmp/google-cloud-sdk.tar.gz \
    && /home/gcloud/google-cloud-sdk/install.sh \
        --usage-reporting false \
        --bash-completion true \
        --rc-path /home/gcloud/.bashrc \
        --path-update true \
    && rm -rf /tmp/google-cloud-sdk.tar.gz

# create an SSH key
RUN ssh-keygen -t rsa -b 2048 -N "" -f /home/gcloud/.ssh/google_compute_engine

#default command
CMD /bin/bash

#example use: 'sudo docker run -it cpswan/gcloud'
