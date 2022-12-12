package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func produce(ctx context.Context, c sarama.Client, topic string) error {
	p, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		return err
	}
	defer p.Close()

	count := 0
	for ctx.Err() == nil {
		partition, offset, err := p.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Testing %d", count)),
		})
		if err != nil {
			return err
		}
		log.Printf("SEND:\tMessage \"Testing %d\", partition: %d, offset: %d\n", count, partition, offset)
		time.Sleep(5 * time.Second)
		count++
	}
	return nil
}
