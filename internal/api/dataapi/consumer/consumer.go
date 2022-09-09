package consumer

import (
	"expvar"
	"sync"

	"gitlab.ozon.dev/krotovkk/homework/internal/counters"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type Consumer struct {
	service ports.Service

	incomeCnt        *counters.Counter
	incomeSuccessCnt *counters.Counter
	incomeFailedCnt  *counters.Counter
}

func NewConsumer(service ports.Service) *Consumer {
	consumer := &Consumer{service: service}
	consumer.incomeCnt = &counters.Counter{M: &sync.RWMutex{}}
	consumer.incomeSuccessCnt = &counters.Counter{M: &sync.RWMutex{}}
	consumer.incomeFailedCnt = &counters.Counter{M: &sync.RWMutex{}}

	expvar.Publish("Income requests counter", consumer.incomeCnt)
	expvar.Publish("Income success requests counter", consumer.incomeSuccessCnt)
	expvar.Publish("Income failed requests counter", consumer.incomeFailedCnt)

	return consumer
}
