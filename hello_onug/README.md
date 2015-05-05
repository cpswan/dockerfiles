##Hello_ONUG

A minimal Docker container with a staticly linked Go binary inside a `FROM scratch` container. Based on the excellent work by [Adriaan de Jonge](https://github.com/adriaandejonge) with his efforts to [Create The Smallest Possible Docker Container](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/).

Compiling static binaries in Go 1.4 has [changed a little](https://github.com/golang/go/issues/9344) so I needed to add the `-installsuffix` flag to the build line:

    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -installsuffix .
