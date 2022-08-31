package brokercommander

import (
	"github.com/Shopify/sarama"
	api "gitlab.ozon.dev/krotovkk/homework/internal/api/dataapi/consumer"
	"gitlab.ozon.dev/krotovkk/homework/internal/common"
)

type brokerRouter struct {
	routes map[string]sarama.ConsumerGroupHandler
}

func NewBrokerRouter() *brokerRouter {
	return &brokerRouter{routes: map[string]sarama.ConsumerGroupHandler{}}
}

func (r *brokerRouter) RegisterRoutes(consumer *api.Consumer) {
	r.routes[common.CartCreate] = &api.CartCreateHandler{Consumer: consumer}
	r.routes[common.CartAddProduct] = &api.AddProductToCartHandler{Consumer: consumer}

	r.routes[common.ProductCreate] = &api.ProductCreateHandler{Consumer: consumer}
	r.routes[common.ProductUpdate] = &api.ProductUpdateHandler{Consumer: consumer}
	r.routes[common.ProductDelete] = &api.ProductDeleteHandler{Consumer: consumer}
}
