package api

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/memorystore"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func NewServer(productService ports.ProductService) pb.AdminServer {
	return &server{
		productService: productService,
	}
}

type server struct {
	pb.UnimplementedAdminServer
	productService ports.ProductService
}

func validateError(err error) error {
	if errors.Is(err, memorystore.ErrProductExist) {
		return status.Error(codes.AlreadyExists, err.Error())
	}

	if errors.Is(err, memorystore.ErrProductNotExist) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}

func (s *server) ProductCreate(ctx context.Context, req *pb.ProductCreateRequest) (*pb.ProductCreateResponse, error) {
	err := s.productService.Create(req.GetName(), req.GetPrice())

	if err != nil {
		return nil, validateError(err)
	}

	return &pb.ProductCreateResponse{}, nil
}

func (s *server) ProductUpdate(ctx context.Context, req *pb.ProductUpdateRequest) (*pb.ProductUpdateResponse, error) {
	err := s.productService.Update(uint(req.GetId()), req.GetName(), req.GetPrice())

	if err != nil {
		return nil, validateError(err)
	}

	return &pb.ProductUpdateResponse{}, nil
}

func (s *server) ProductList(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {
	products := s.productService.List()
	result := make([]*pb.ProductListResponse_Product, 0, len(products))
	for _, product := range products {
		result = append(result, &pb.ProductListResponse_Product{
			Id:    uint64(product.GetId()),
			Name:  product.GetName(),
			Price: product.GetPrice(),
		})
	}

	return &pb.ProductListResponse{Products: result}, nil
}

func (s *server) ProductDelete(ctx context.Context, req *pb.ProductDeleteRequest) (*pb.ProductDeleteResponse, error) {
	err := s.productService.Delete(uint(req.GetId()))

	if err != nil {
		return nil, validateError(err)
	}

	return &pb.ProductDeleteResponse{}, nil
}
