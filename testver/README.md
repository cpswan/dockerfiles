testver
=======

What happens when I update something on Docker Hub?

This is a trivial Dockerfile to test out the linkage between a local registry and Docker Hub.

`sudo docker run -it cpswan/testver` will pull down the :latest version from Docker Hub when it is first run.

Next time I run `sudo docker run -it cpswan/testver` (or even `cpswan/testver:latest`) docker will look in my local registry and find the image there.

If Docker Hub has been updated with a newer version then I have to explicitly `docker pull cpswan/testver` in order to get it.
