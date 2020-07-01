Get apache pulsar docker image
`
sudo docker run -it \
  -p 6650:6650 \
  -p 8080:8080 \
  -v $PWD/data:/pulsar/data \
  apachepulsar/pulsar:2.4.0 \
  bin/pulsar standalone
 `
go get -u "github.com/apache/pulsar-client-go/pulsar"
