package broker

import (
	"encoding/json"
	"expvar"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/cache/rediscache"
	"gitlab.ozon.dev/krotovkk/homework/internal/counters"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	"go.opencensus.io/trace"
)

type brokerService struct {
	cartProducer    ports.CartService
	productProducer ports.ProductService
	client          *redis.Client

	producer sarama.SyncProducer
	outCnt   *counters.Counter
}

func NewBrokerService(producer sarama.SyncProducer, client *redis.Client) ports.Service {
	service := &brokerService{producer: producer}
	service.outCnt = &counters.Counter{M: &sync.RWMutex{}}
	service.cartProducer = NewCartService(service)
	service.productProducer = NewProductService(service)
	service.client = client

	expvar.Publish("Out requests counter", service.outCnt)

	return service
}

func (bs *brokerService) Product() ports.ProductService {
	return bs.productProducer
}

func (bs *brokerService) Cart() ports.CartService {
	return bs.cartProducer
}

func (bs *brokerService) sendTrace(span *trace.Span, topic string) error {
	spanJson, err := json.Marshal(span)
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "send span", "error": err}).Warnf("Fail while span marshaling")
		return err
	}

	_, _, err = bs.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(0),
		Value: sarama.ByteEncoder(spanJson),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "send span", "error": err}).Warnf("Fail to send messege to broker")
		return err
	}

	return err
}

func (bs *brokerService) readProductsFromCache(chanel string) ([]*model.Product, error) {
	subscriber := bs.client.Subscribe(chanel)

	for {
		logrus.WithFields(logrus.Fields{"chanel": chanel}).Infof("products reading...")
		msg, err := subscriber.ReceiveMessage()
		if err != nil {
			return nil, err
		}

		var products rediscache.ProductsData

		err = json.Unmarshal([]byte(msg.Payload), &products)
		if err != nil {
			return nil, err
		}

		logrus.WithFields(logrus.Fields{"chanel": chanel, "products": products}).Infof("products readed")

		return products, nil
	}
}
