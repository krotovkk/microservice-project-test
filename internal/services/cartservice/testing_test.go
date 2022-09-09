package cartservice

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	mock_stores "gitlab.ozon.dev/krotovkk/homework/internal/store/mocks"
)

type cartServiceFixture struct {
	cartService ports.CartService
	cartStore   *mock_stores.MockCartStore
	ctrl        *gomock.Controller
}

func setUpFixture(t *testing.T) *cartServiceFixture {
	ctrl := gomock.NewController(t)
	cartStore := mock_stores.NewMockCartStore(ctrl)
	cartService := NewCartService(&Options{Store: cartStore})

	return &cartServiceFixture{cartService: cartService, cartStore: cartStore, ctrl: ctrl}
}

func (f *cartServiceFixture) tearDown() {
	f.ctrl.Finish()
}
