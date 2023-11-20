package service

import (
	"context"

	pb "github.com/thetnaingtn/go-boutique/api/product/v1"
	"github.com/thetnaingtn/go-boutique/app/product/internal/biz"
)

type ProductService struct {
	pb.UnimplementedProductServer
	uc *biz.ProductUsecase
}

func NewProductService(uc *biz.ProductUsecase) *ProductService {
	return &ProductService{uc: uc}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {

	product, err := s.uc.CreateProduct(ctx, &biz.Product{
		Name:        req.Body.GetName(),
		Description: req.Body.GetDescription(),
		Price:       req.Body.GetPrice(),
	})

	if err != nil {
		return &pb.CreateProductReply{}, err
	}

	reply := &pb.CreateProductReply{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	return reply, nil
}
func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {

	up := &biz.UpdateProduct{
		Name:        req.Body.Name,
		Price:       req.Body.Price,
		Description: req.Body.Description,
	}

	p, err := s.uc.Update(ctx, req.GetId(), up)

	if err != nil {
		return &pb.UpdateProductReply{}, err
	}

	updatedProduct := &pb.UpdateProductReply{
		Id:          req.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	return updatedProduct, nil
}
func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {

	id, err := s.uc.DeleteById(ctx, req.Id)

	if err != nil {
		return &pb.DeleteProductReply{}, err
	}

	return &pb.DeleteProductReply{Id: id}, nil
}
func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	p, err := s.uc.GetProduct(ctx, req.GetId())

	if err != nil {
		return &pb.GetProductReply{}, err
	}

	product := &pb.GetProductReply{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	return product, nil
}
func (s *ProductService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductReply, error) {
	return &pb.ListProductReply{}, nil
}
