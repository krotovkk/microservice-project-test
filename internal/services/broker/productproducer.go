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

type ProductProducer struct {
	*brokerService
}

func NewProductService(service *brokerService) *ProductProducer {
	return &ProductProducer{brokerService: service}
}

func (ps *ProductProducer) CreateProduct(ctx context.Context, name string, price float64) (*model.Product, error) {
	_, span := trace.StartSpan(ctx, "Send create product messege to broker")
	defer span.End()

	product := model.Product{Name: name, Price: price}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "create product", "error": err}).Warnf("Fail while marshaling")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{"operation": "update product"}).Infof("send messege to broker")
	_, _, err = ps.producer.SendMessage(&sarama.ProducerMessage{
		Topic: common.ProductCreate,
		Key:   sarama.StringEncoder(0),
		Value: sarama.ByteEncoder(jsonProduct),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "create product", "error": err}).Warnf("Fail to send messege to broker")
		return nil, err
	}
	logrus.WithFields(logrus.Fields{"operation": "create product"}).Infof("Successfuly sended")

	return nil, err
}

func (ps *ProductProducer) UpdateProduct(ctx context.Context, name string, price float64, id uint) error {
	_, span := trace.StartSpan(ctx, "Send update product messege to broker")
	defer span.End()

	product := model.Product{Name: name, Price: price, Id: id}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "update product", "error": err}).Warnf("Fail while marshaling")
		return err
	}

	logrus.WithFields(logrus.Fields{"operation": "update product"}).Infof("send messege to broker")
	_, _, err = ps.producer.SendMessage(&sarama.ProducerMessage{
		Topic: common.ProductUpdate,
		Key:   sarama.StringEncoder(id),
		Value: sarama.ByteEncoder(jsonProduct),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "update product", "error": err}).Warnf("Fail to send messege to broker")
		return err
	}
	logrus.WithFields(logrus.Fields{"operation": "update product"}).Infof("Successfuly sended")

	return err
}

func (ps *ProductProducer) DeleteProduct(ctx context.Context, id uint) error {
	_, span := trace.StartSpan(ctx, "Send update product messege to broker")
	defer span.End()

	product := model.Product{Id: id}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "delete product", "error": err}).Warnf("Fail while marshaling")
		return err
	}

	logrus.WithFields(logrus.Fields{"operation": "delete product"}).Infof("send messege to broker")
	_, _, err = ps.producer.SendMessage(&sarama.ProducerMessage{
		Topic: common.ProductDelete,
		Key:   sarama.StringEncoder(id),
		Value: sarama.ByteEncoder(jsonProduct),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"operation": "delete product", "error": err}).Warnf("Fail to send messege to broker")
		return err
	}
	logrus.WithFields(logrus.Fields{"operation": "delete product"}).Infof("Successfuly sended")

	return err
}

func (ps *ProductProducer) GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error) {
	return nil, nil
}
