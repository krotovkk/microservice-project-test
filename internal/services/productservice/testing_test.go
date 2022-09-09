package productservice

import (
	"testing"

	"github.com/golang/mock/gomock"

	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	mock_stores "gitlab.ozon.dev/krotovkk/homework/internal/store/mocks"
)

type productServiceFixture struct {
	productService ports.ProductService
	productStore   *mock_stores.MockProductStore
	ctrl           *gomock.Controller
}

func setUpFixture(t *testing.T) *productServiceFixture {
	ctrl := gomock.NewController(t)
	productStore := mock_stores.NewMockProductStore(ctrl)
	productService := NewProductService(&Options{Store: productStore})

	return &productServiceFixture{productService: productService, productStore: productStore, ctrl: ctrl}
}

func (f *productServiceFixture) tearDown() {
	f.ctrl.Finish()
}
