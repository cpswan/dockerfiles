# Getting Started with WebVR

Dev environment for WebVR based on [A-Frame](https://github.com/aframevr/aframe) using Ada Rose Edwards' [getting-started-with-webvr](https://github.com/AdaRoseEdwards/getting-started-with-webvr) samples and examples.

Installs [GitHub Pages Ruby Gem](https://github.com/github/pages-gem) and dependencies using [bundler](http://bundler.io/) and provides te JavaScript runtime needed by extjs by installing the nodejs package.

### Example usage:

First start a container that will run the GitHub pages Jekyll environment:

`sudo docker run -dp 4000:4000 --name webvr cpswan/webvr`

It's now possible to browse to Ada's samples and examples at http://dockerhostname:4000

To modify the samples and examples, or add files of your own open a shell to the container:

`sudo docker exec -it webvr bash`

and then use your preferred editor.
