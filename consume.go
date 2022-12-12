package main

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

func consume(ctx context.Context, c sarama.Client, topic string) error {
	g, err := sarama.NewConsumerGroupFromClient("sample-group", c)
	if err != nil {
		return err
	}
	defer g.Close()

	consumer := consumer{}
	for {
		topics := []string{topic}
		err := g.Consume(ctx, topics, consumer)
		if err != nil {
			return err
		}
		if ctx.Err() != nil {
			return nil
		}
	}
}

type consumer struct{}

func (consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("RECEIVE:\tMessage %q, partition: %d, offset: %d\n", msg.Value, msg.Partition, msg.Offset)
		session.MarkMessage(msg, "")
	}
	return nil
}

func (consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
