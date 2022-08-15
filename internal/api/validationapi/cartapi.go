package validationapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

type cartServer struct {
	pb.UnimplementedCartServer
	client pb.CartClient
}

func NewCartServer(client pb.CartClient) *cartServer {
	return &cartServer{client: client}
}

func (s *cartServer) CartCreate(ctx context.Context, req *pb.CartCreateRequest) (*pb.CartCreateResponse, error) {
	return s.client.CartCreate(ctx, req)
}

func (s *cartServer) CartGetProducts(ctx context.Context, req *pb.CartGetProductsRequest) (*pb.CartGetProductsResponse, error) {
	cart := model.Cart{Id: req.GetId()}

	if err := cart.CheckId(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.client.CartGetProducts(ctx, req)
}

func (s *cartServer) CartAddProduct(ctx context.Context, req *pb.CartAddProductRequest) (*pb.CartAddProductResponse, error) {
	product := model.Product{Id: uint(req.GetProductId())}
	if err := product.CheckId(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	cart := model.Cart{Id: req.GetCartId()}

	if err := cart.CheckId(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.client.CartAddProduct(ctx, req)
}
