package dataapi

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/memorystore"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

const streamProductsLimit = 5

func NewProductServer(productService ports.ProductService) pb.ProductServer {
	return &productServer{
		productService: productService,
	}
}

type productServer struct {
	pb.UnimplementedProductServer
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

func (s *productServer) ProductCreate(ctx context.Context, req *pb.ProductCreateRequest) (*pb.ProductCreateResponse, error) {
	p, err := s.productService.CreateProduct(ctx, req.GetName(), req.GetPrice())

	if err != nil {
		return nil, validateError(err)
	}

	return &pb.ProductCreateResponse{Id: int64(p.Id), Name: p.Name, Price: p.Price}, nil
}

func (s *productServer) ProductUpdate(ctx context.Context, req *pb.ProductUpdateRequest) (*pb.ProductUpdateResponse, error) {
	err := s.productService.UpdateProduct(ctx, req.GetName(), req.GetPrice(), uint(req.GetId()))

	if err != nil {
		return nil, validateError(err)
	}

	return &pb.ProductUpdateResponse{}, nil
}

func (s *productServer) ProductList(req *pb.ProductListRequest, res pb.Product_ProductListServer) error {
	products, err := s.productService.GetAllProducts(context.Background(), req.GetLimit(), req.GetOffset())
	if err != nil {
		return err
	}

	buffer := make([]*pb.ProductListResponse_Product, 0, streamProductsLimit)
	for index, product := range products {
		buffer = append(buffer, &pb.ProductListResponse_Product{
			Id:    uint64(product.GetId()),
			Name:  product.GetName(),
			Price: product.GetPrice(),
		})
		if len(buffer) == streamProductsLimit || index == len(products)-1 {
			res.Send(&pb.ProductListResponse{Products: buffer})
			buffer = buffer[:0]
		}
	}

	return nil
}

func (s *productServer) ProductDelete(ctx context.Context, req *pb.ProductDeleteRequest) (*pb.ProductDeleteResponse, error) {
	err := s.productService.DeleteProduct(ctx, uint(req.GetId()))

	if err != nil {
		return nil, validateError(err)
	}

	return &pb.ProductDeleteResponse{}, nil
}
