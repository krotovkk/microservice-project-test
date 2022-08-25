package dataapi

import (
	"context"

	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

type cartServer struct {
	pb.UnimplementedCartServer

	service ports.CartService
}

func NewCartServer(service ports.CartService) *cartServer {
	return &cartServer{service: service}
}

func (s *cartServer) CartCreate(ctx context.Context, req *pb.CartCreateRequest) (*pb.CartCreateResponse, error) {
	c, err := s.service.CreateCart(ctx)
	return &pb.CartCreateResponse{Id: c.Id, CreatedAt: c.CreatedAt}, err
}

func (s *cartServer) CartGetProducts(ctx context.Context, req *pb.CartGetProductsRequest) (*pb.CartGetProductsResponse, error) {
	products, err := s.service.GetCartProducts(ctx, req.GetId())

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
	err := s.service.AddProductToCart(ctx, req.GetProductId(), req.GetCartId())

	if err != nil {
		return nil, err
	}

	return &pb.CartAddProductResponse{}, nil
}
