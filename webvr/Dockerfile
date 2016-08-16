FROM ruby

MAINTAINER @cpswan

# install utility packages and dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
  sudo curl wget vim bundler zlib1g-dev libxml2-dev nodejs

# add an webvr user and allow it to sudo
RUN useradd -m webvr &&\
  echo 'webvr ALL=NOPASSWD: ALL' > /etc/sudoers.d/webvr

USER webvr

# install getting-started-with-webvr repo
RUN cd /home/webvr &&\
  git clone https://github.com/AdaRoseEdwards/getting-started-with-webvr.git &&\
  cd getting-started-with-webvr &&\
  bundle config build.nokogiri --use-system-libraries &&\
  bundle install &&\
  echo 'host: 0.0.0.0' >> _config.yml

#default command
WORKDIR /home/webvr/getting-started-with-webvr
CMD bundle exec jekyll serve 

#example usage: 'sudo docker run -dp 4000:4000 --name webvr cpswan/webvr'
#example usage: 'sudo docker exec -it webvr bash'
