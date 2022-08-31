package consumer

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/broker"
	"go.opencensus.io/trace"
)

type CartCreateHandler struct {
	*Consumer
}

func (pc *CartCreateHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *CartCreateHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *CartCreateHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()

	for msg := range claim.Messages() {
		pc.incomeCnt.Inc()
		session.MarkMessage(msg, "")
		var spanContext trace.SpanContext
		err := json.Unmarshal(msg.Value, &spanContext)

		if err != nil {
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Fail to unmarshal")
		}

		_, span := trace.StartSpanWithRemoteParent(context.Background(), "create cart", spanContext)
		defer span.End()

		logrus.WithFields(logrus.Fields{"topic": msg.Topic, "traceId": span.SpanContext().TraceID.String()}).Infof("Span info")

		c, err := pc.service.Cart().CreateCart(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Failed operation")
			pc.incomeFailedCnt.Inc()
			continue
		}
		pc.incomeSuccessCnt.Inc()
		logrus.WithFields(logrus.Fields{"topic": msg.Topic, "cart": c}).Infof("Successfuly operation")
	}

	return nil
}

type AddProductToCartHandler struct {
	*Consumer
}

func (pc *AddProductToCartHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *AddProductToCartHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *AddProductToCartHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()

	for msg := range claim.Messages() {
		pc.incomeCnt.Inc()
		session.MarkMessage(msg, "")
		ids := &broker.CartProductIdsWithSpan{}

		err := json.Unmarshal(msg.Value, ids)
		if err != nil {
			pc.incomeFailedCnt.Inc()
			logrus.WithFields(logrus.Fields{"topic": msg.Topic, "error": err}).Warnf("Fail to unmarshal")
			continue
		}

		_, span := trace.StartSpanWithRemoteParent(context.Background(), "create cart", ids.SpanContext)
		defer span.End()
		logrus.WithFields(logrus.Fields{"topic": msg.Topic, "traceId": span.SpanContext().TraceID.String()}).Infof("Span info")

		err = pc.service.Cart().AddProductToCart(ctx, ids.ProductId, ids.CartId)
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
