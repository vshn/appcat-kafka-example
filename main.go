package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	topic := flag.String("topic", "test", "What topic to send to or receive from. The topic will be created if it does not exist")
	uri := flag.String("uri", "", "The URI of the Kafka Cluster")
	cert := flag.String("cert", "", "Path to the client certificate")
	key := flag.String("key", "", "Path to the client key")
	ca := flag.String("ca", "", "The Kafka's CA certificate")
	flag.Parse()

	if *uri == "" {
		*uri = os.Getenv("KAFKA_URI")
	}

	c, err := kafkaClient(*uri, *cert, *key, *ca)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	err = createTopic(c, *topic)
	if err != nil {
		log.Fatalf("Failed to create topic: %s", err.Error())
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		err = consume(ctx, c, *topic)
		if err != nil {
			log.Fatalf("Failed to consume: %s", err.Error())
		}
	}()
	err = produce(ctx, c, *topic)
	if err != nil {
		log.Fatalf("Failed to produce: %s", err.Error())
	}
}

func createTopic(c sarama.Client, topic string) error {
	topics, err := c.Topics()
	if err != nil {
		log.Fatalf("topic list: %s", err.Error())
	}
	for _, t := range topics {
		if t == topic {
			return nil
		}
	}

	ctrl, err := c.Controller()
	if err != nil {
		return err
	}
	_, err = ctrl.CreateTopics(&sarama.CreateTopicsRequest{
		Timeout: time.Second,
		TopicDetails: map[string]*sarama.TopicDetail{topic: {
			NumPartitions:     1,
			ReplicationFactor: 1,
		}},
	})
	return err

}

func kafkaClient(uri, certPath, keyPath, caPath string) (sarama.Client, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_2_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Net.TLS.Enable = true

	caCertPool := x509.NewCertPool()
	caCert, err := os.ReadFile(caPath)
	if err != nil {
		return nil, err
	}
	caCertPool.AppendCertsFromPEM([]byte(caCert))

	keypair, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	config.Net.TLS.Config = &tls.Config{
		Certificates: []tls.Certificate{keypair},
		RootCAs:      caCertPool,
	}

	return sarama.NewClient([]string{uri}, config)
}
