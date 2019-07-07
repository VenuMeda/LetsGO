package kafka

import (
	"fmt"
)

/*
I have seen far by standing on the shoulders of giants.
Example implemented in top of [cluster "github.com/bsm/sarama-cluster"] and ["github.com/segmentio/kafka-go"]
*/
import (
	cluster "github.com/bsm/sarama-cluster"
)

/*
Kafka subscriber that once we receive the message
*/
func SubscribeConsumer(broker Broker, topic Topic, fn func(str string)) {
	consumer, _ := createConsumer(broker.Value, topic.Value)
	for consumerMessage := range consumer.Messages() {
		switch consumerMessage.Topic {
		case topic.Value:
			updatedEvent := string(consumerMessage.Value) + " and Kafka consumer"
			fmt.Printf("KAFKA Consumer: %s \n", updatedEvent)
			fn(updatedEvent)
			break
		default:
			println("Error kafka message not expected for any Topic")
		}
	}
}

func createConsumer(broker string, topic string) (*cluster.Consumer, error) {
	config := cluster.NewConfig()
	return cluster.NewConsumer(
		[]string{broker},
		"group-id",
		[]string{topic},
		config)
}