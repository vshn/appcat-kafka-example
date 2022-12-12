# AppCat Kafka Sample

This is an example application to connect to a Kafka instance.
It uses client side certificates to authenticate, will create a test topic, and continuously publish to and read from this topic.


## Build

You can easily build the example using make.

    make build-bin

To build the docker container you can run.

    CONTAINER_IMG="foo/bar:tag" make build-docker


## Run

To run the example producer and consumer you need access to a Kafka instance, as well as the client certificate and key, and the CA of the Kafka instance.
This information can be passed as command line flags.

    ./appcat-kafka-example --topic foobar --uri my-kafka.example.com:21701 --ca ca.crt --cert service.cert --key service.key   
