package validationapi

import (
	"context"

	"gitlab.ozon.dev/krotovkk/homework/internal/ports"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

type cartServer struct {
	pb.UnimplementedCartServer
	client  pb.CartClient
	service ports.CartService
}

func NewCartServer(client pb.CartClient, service ports.CartService) *cartServer {
	return &cartServer{
		client:  client,
		service: service,
	}
}

func (s *cartServer) CartCreate(ctx context.Context, req *pb.CartCreateRequest) (*pb.CartCreateResponse, error) {
	_, err := s.service.CreateCart(ctx)

	return &pb.CartCreateResponse{}, err
}

func (s *cartServer) CartGetProducts(ctx context.Context, req *pb.CartGetProductsRequest) (*pb.CartGetProductsResponse, error) {
	cart := model.Cart{Id: req.GetId()}

	if err := cart.CheckId(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	products, err := s.service.GetCartProducts(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	result := make([]*pb.CartGetProductsResponse_Product, 0, len(products))

	for _, product := range products {
		result = append(result, &pb.CartGetProductsResponse_Product{
			Id:    uint64(product.Id),
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return &pb.CartGetProductsResponse{Products: result}, nil
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

	err := s.service.AddProductToCart(ctx, int64(product.Id), cart.Id)

	return &pb.CartAddProductResponse{}, err
}
