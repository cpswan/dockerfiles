awscli
======

### CLI for Amazon Web Services

Inpired by [Martin Elwin's](https://twitter.com/elvvin) presentation [Deep Dive: AWS Command Line Interface](http://www.slideshare.net/AmazonWebServices/deep-dive-advanced-usage-of-the-aws-cli) at the AWS Summit in London 2015 this includes jpterm, the [JMESPath Terminal](https://github.com/jmespath/jmespath.terminal).

Example usage:

`sudo docker run -it cpswan/awscli`

Once the container has downloaded and started use `aws configure` to set up default keys etc. With the keys in place `aws ec2 describe-regions` can be used to test that everything is working.
