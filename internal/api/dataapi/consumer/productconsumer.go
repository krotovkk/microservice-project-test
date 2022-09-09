package consumer

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

type ProductCreateHandler struct {
	*Consumer
}

func (pc *ProductCreateHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductCreateHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductCreateHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()
	for msg := range claim.Messages() {
		pc.incomeCnt.Inc()
		session.MarkMessage(msg, "")
		product := &model.Product{}
		err := json.Unmarshal(msg.Value, product)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Fail to unmarshal")
			continue
		}

		_, err = pc.service.Product().CreateProduct(ctx, product.Name, product.Price)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Failed operation")
			continue
		}
		pc.incomeSuccessCnt.Inc()
		logrus.WithFields(logrus.Fields{"topic": msg.Topic}).Infof("Successfuly operation")
	}

	return nil
}

type ProductUpdateHandler struct {
	*Consumer
}

func (pc *ProductUpdateHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductUpdateHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductUpdateHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()

	for msg := range claim.Messages() {
		pc.incomeCnt.Inc()
		session.MarkMessage(msg, "")
		product := &model.Product{}
		err := json.Unmarshal(msg.Value, product)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Fail to unmarshal")
			continue
		}

		err = pc.service.Product().UpdateProduct(ctx, product.Name, product.Price, product.Id)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Failed operation")
			continue
		}
		pc.incomeSuccessCnt.Inc()
		logrus.WithFields(logrus.Fields{"topic": msg.Topic}).Infof("Successfuly operation")
	}

	return nil
}

type ProductDeleteHandler struct {
	*Consumer
}

func (pc *ProductDeleteHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductDeleteHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductDeleteHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()

	for msg := range claim.Messages() {
		pc.incomeCnt.Inc()
		session.MarkMessage(msg, "")
		product := &model.Product{}
		err := json.Unmarshal(msg.Value, product)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Fail to unmarshal")
			continue
		}

		err = pc.service.Product().DeleteProduct(ctx, product.Id)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Failed operation")
			continue
		}
		pc.incomeSuccessCnt.Inc()
		logrus.WithFields(logrus.Fields{"topic": msg.Topic}).Infof("Successfuly operation")
	}

	return nil
}

type ProductListHandler struct {
	*Consumer
}

func (pc *ProductListHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductListHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *ProductListHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()

	for msg := range claim.Messages() {
		pc.incomeCnt.Inc()
		session.MarkMessage(msg, "")
		limitOffset := &model.LimitOffset{}
		err := json.Unmarshal(msg.Value, limitOffset)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Fail to unmarshal")
			continue
		}

		_, err = pc.service.Product().GetAllProducts(ctx, limitOffset.Limit, limitOffset.Offset)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Failed operation")
			continue
		}
		pc.incomeSuccessCnt.Inc()
		logrus.WithFields(logrus.Fields{"topic": msg.Topic}).Infof("Successfuly operation")
	}

	return nil
}
