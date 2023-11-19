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

type ProductRepo interface {
	Save(context.Context, *Product) (*Product, error)
	Update(context.Context, *Product) (*Product, error)
	FindByID(context.Context, int64) (*Product, error)
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
	// uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
