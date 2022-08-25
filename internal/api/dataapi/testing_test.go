package dataapi

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_services "gitlab.ozon.dev/krotovkk/homework/internal/services/mocks"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

type productServerFixture struct {
	ctrl    *gomock.Controller
	service *mock_services.MockProductService
	server  pb.ProductServer
}

func setUpProductFixture(t *testing.T) *productServerFixture {
	ctrl := gomock.NewController(t)
	service := mock_services.NewMockProductService(ctrl)
	server := NewProductServer(service)

	return &productServerFixture{ctrl: ctrl, service: service, server: server}
}

func (f *productServerFixture) tearDown() {
	f.ctrl.Finish()
}

type cartServerFixture struct {
	ctrl    *gomock.Controller
	service *mock_services.MockCartService
	server  pb.CartServer
}

func setUpCartFixture(t *testing.T) *cartServerFixture {
	ctrl := gomock.NewController(t)
	service := mock_services.NewMockCartService(ctrl)
	server := NewCartServer(service)

	return &cartServerFixture{ctrl: ctrl, service: service, server: server}
}

func (f *cartServerFixture) tearDown() {
	f.ctrl.Finish()
}
