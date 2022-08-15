package validationapi

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

var (
	errInvalidOffset = errors.New("wrong offset value, must be greater than 0")
	errInvalidLimit  = errors.New("wrong limit value, must be greater than 0")
)

type productServer struct {
	pb.UnimplementedProductServer
	client pb.ProductClient
}

func NewProductServer(productClient pb.ProductClient) pb.ProductServer {
	return &productServer{
		client: productClient,
	}
}

func (s *productServer) ProductCreate(ctx context.Context, req *pb.ProductCreateRequest) (*pb.ProductCreateResponse, error) {
	product := model.Product{Price: req.GetPrice(), Name: req.GetName()}

	if err := product.CheckPrice(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := product.CheckName(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.client.ProductCreate(ctx, req)
}

func (s *productServer) ProductUpdate(ctx context.Context, req *pb.ProductUpdateRequest) (*pb.ProductUpdateResponse, error) {
	product := model.Product{Id: uint(req.GetId()), Price: req.GetPrice(), Name: req.GetName()}

	if err := product.CheckId(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := product.CheckPrice(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := product.CheckName(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.client.ProductUpdate(ctx, req)
}

func (s *productServer) ProductList(req *pb.ProductListRequest, res pb.Product_ProductListServer) error {
	if req.GetOffset() < 0 {
		return status.Error(codes.InvalidArgument, errInvalidOffset.Error())
	}

	if req.GetLimit() < 0 {
		return status.Error(codes.InvalidArgument, errInvalidLimit.Error())
	}

	response, err := s.client.ProductList(context.Background(), req)

	if err != nil {
		return err
	}

	for {
		products, err := response.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		res.Send(products)
	}

	return nil
}

func (s *productServer) ProductDelete(ctx context.Context, req *pb.ProductDeleteRequest) (*pb.ProductDeleteResponse, error) {
	product := model.Product{Id: uint(req.GetId())}

	if err := product.CheckId(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.client.ProductDelete(ctx, req)
}
