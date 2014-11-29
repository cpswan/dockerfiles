FROM ubuntu:14.04

MAINTAINER cpswan

RUN export DEBIAN_FRONTEND=noninteractive &&\
    apt-get update &&\
    apt-get install -y build-essential subversion git-core libncurses5-dev \
      zlib1g-dev gawk flex quilt libssl-dev xsltproc libxml-parser-perl \
      mercurial bzr ecj cvs unzip python wget &&\
    useradd -m openwrt &&\
    echo 'openwrt ALL=NOPASSWD: ALL' > /etc/sudoers.d/openwrt &&\
    wget http://d.wrtnode.com/3.14.18/Wrtnode_SDK_11_27.tar.bz2 &&\
    wget http://d.wrtnode.com/3.14.18/dl_11_27.tar.bz2 &&\
    sudo -iu openwrt sudo tar -xvf /Wrtnode_SDK_11_27.tar.bz2 -C \
      /home/openwrt/ &&\
    sudo -iu openwrt sudo tar -xvf /dl_11_27.tar.bz2 -C \
      /home/openwrt/wrtnode-sdk/ &&\
    rm /Wrtnode_SDK_11_27.tar.bz2 &&\
    rm /dl_11_27.tar.bz2 

CMD /usr/bin/sudo -iu openwrt /bin/bash

