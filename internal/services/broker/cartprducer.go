package broker

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/common"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"go.opencensus.io/trace"
)

type CartProductIdsWithSpan struct {
	CartId    int64 `json:"cart_id"`
	ProductId int64 `json:"product_id"`
	trace.SpanContext
}

type CartProducer struct {
	*brokerService
}

func NewCartService(service *brokerService) *CartProducer {
	return &CartProducer{brokerService: service}
}

func (cs *CartProducer) CreateCart(ctx context.Context) (*model.Cart, error) {
	_, span := trace.StartSpan(ctx, "Send create cart messege to broker")
	defer span.End()

	spanJson, err := json.Marshal(span.SpanContext())
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "send span", "error": err}).Warnf("Fail while span marshaling")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{"operation": "create cart", "span": span.SpanContext().TraceID.String()}).Infof("Send messege to broker")
	_, _, err = cs.producer.SendMessage(&sarama.ProducerMessage{
		Topic: common.CartCreate,
		Key:   sarama.StringEncoder(0),
		Value: sarama.ByteEncoder(spanJson),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "create cart"}).Warnf("Fail to send messege to broker")
		return nil, err
	}
	logrus.WithFields(logrus.Fields{"operation": "create cart", "error": err}).Infof("Successfuly sended")

	return nil, err
}

func (cs *CartProducer) GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error) {
	return nil, nil
}

func (cs *CartProducer) AddProductToCart(ctx context.Context, productId, cartId int64) error {
	_, span := trace.StartSpan(ctx, "Send add product to cart messege to broker")
	defer span.End()

	ids := &CartProductIdsWithSpan{
		CartId:      cartId,
		ProductId:   productId,
		SpanContext: span.SpanContext(),
	}

	jsonValue, err := json.Marshal(ids)
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "add product to cart", "error": err}).Warnf("Fail while marshaling")
		return err
	}

	logrus.WithFields(logrus.Fields{"operation": "add product to cart"}).Infof("send messege to broker")
	_, _, err = cs.producer.SendMessage(&sarama.ProducerMessage{
		Topic: common.CartAddProduct,
		Key:   sarama.StringEncoder(cartId),
		Value: sarama.ByteEncoder(jsonValue),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "add product to cart"}).Warnf("Fail to send messege to broker")
		return err
	}
	logrus.WithFields(logrus.Fields{"operation": "add product to cart"}).Infof("Successfuly sended")

	return err
}
