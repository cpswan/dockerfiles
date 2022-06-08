## Hello_ver

A minimal Docker container with a staticly linked Go binary inside a
`FROM scratch` container. Based on the excellent work by
[Adriaan de Jonge](https://github.com/adriaandejonge) with his efforts to
[Create The Smallest Possible Docker Container](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/).

### Purpose

This Dockerfile and its contents are used to create multiple versions of an
image to test the addition of `DONT_WAIT` to the atsign-company fork of
[Shepherd](https://github.com/atsign-company/shepherd).

### Building

Compile the binary:

```bash
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -installsuffix .
```

Build the image:

```bash
sudo docker build -t cpswan/hello_ver .
```

Push the image:

```bash
sudo docker push cpswan/hello_ver
```

### Testing

Spin up a bunch of services using the image:

```bash
for i in {9000..9010}; do sudo docker service create -p $i:$i --constraint node.role==manager cpswan/hello_ver /hello_ver $i; done
```

Bump the version, build a new binary, build a new image, then run Shepherd
without DONT_WAIT to measure speedup:

```bash
time FILTER_SERVICES="" RUN_ONCE_AND_EXIT="true" ./shepherd
```

Bump the version, build a new binary, build a new image, then run Shepherd
with DONT_WAIT to measure speedup:

```bash
time FILTER_SERVICES="" DONT_WAIT="true" RUN_ONCE_AND_EXIT="true" ./shepherd
```