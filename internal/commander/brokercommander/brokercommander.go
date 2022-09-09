package brokercommander

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/config"
	api "gitlab.ozon.dev/krotovkk/homework/internal/api/dataapi/consumer"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type brokerCommander struct {
	router *brokerRouter
}

func NewBrokerCommander() *brokerCommander {
	return &brokerCommander{
		router: NewBrokerRouter(),
	}
}

func Run(service ports.Service, ch chan struct{}) {
	defer func() { ch <- struct{}{} }()

	cfg := sarama.NewConfig()

	ctx := context.Background()
	brokerCommander := NewBrokerCommander()

	consumer := api.NewConsumer(service)
	brokerCommander.router.RegisterRoutes(consumer)

	for topic, handler := range brokerCommander.router.routes {
		go func(topic string, handler sarama.ConsumerGroupHandler) {
			client, err := sarama.NewConsumerGroup(config.Brokers, fmt.Sprintf("%s", topic), cfg)
			if err != nil {
				logrus.WithFields(logrus.Fields{"brokers": config.Brokers, "error": err}).Fatalf("Can't connect to brockers")
			}
			for {
				if err := client.Consume(ctx, []string{topic}, handler); err != nil {
					logrus.WithFields(logrus.Fields{"topic": topic, "error": err}).Warnf("On consume")
				}
			}
		}(topic, handler)
	}
}
