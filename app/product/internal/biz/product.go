package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Product struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
}

type UpdateProduct struct {
	Name        *string `db:"name"`
	Description *string `db:"description"`
	Price       *int64  `db:"price"`
}

type ProductRepo interface {
	Save(context.Context, *Product) (*Product, error)
	Update(context.Context, string, *UpdateProduct) (*Product, error)
	FindByID(context.Context, string) (*Product, error)
	ListByHello(context.Context, string) ([]*Product, error)
	ListAll(context.Context) ([]*Product, error)
}

type ProductUsecase struct {
	repo ProductRepo
	log  *log.Helper
}

func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, g *Product) (*Product, error) {
	return uc.repo.Save(ctx, g)
}

func (uc *ProductUsecase) GetProduct(ctx context.Context, id string) (*Product, error) {
	return uc.repo.FindByID(ctx, id)
}
