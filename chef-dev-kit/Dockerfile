FROM centos:centos6
MAINTAINER cpswan

RUN yum -y install tree unzip
RUN curl -LO https://opscode-omnibus-packages.s3.amazonaws.com/el/6/x86_64/chefdk-0.2.0-2.el6.x86_64.rpm
RUN rpm -i chefdk-0.2.0-2.el6.x86_64.rpm

CMD ["/bin/bash"]
