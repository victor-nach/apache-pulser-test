## Get apache pulsar docker image
`
sudo docker run -it \
  -p 6650:6650 \
  -p 8080:8080 \
  -v $PWD/data:/pulsar/data \
  apachepulsar/pulsar:2.4.0 \
  bin/pulsar standalone
 `
## Get pulsar client for Go
`go get -u "github.com/apache/pulsar-client-go/pulsar"`

## Run a docker image from a docker file
- create dockerfile
- run `docker build -t <image-name> <dockerfile-location>` to build new image
 - e.g `docker build -t producer .`
- run `docker images` to see a list of all images, you should see the image you just created
- run `docker run `
- run `docker rmi <image-name1> <image-name2>` to remove images
- run `docker container prune` to remove stopped containers




