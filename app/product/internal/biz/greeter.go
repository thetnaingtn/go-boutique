package biz

import (
	"context"

	v1 "github.com/thetnaingtn/go-boutique/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

type Product struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
}

// GreeterRepo is a Greater repo.
type ProductRepo interface {
	Save(context.Context, *Product) (*Product, error)
	Update(context.Context, *Product) (*Product, error)
	FindByID(context.Context, int64) (*Product, error)
	ListByHello(context.Context, string) ([]*Product, error)
	ListAll(context.Context) ([]*Product, error)
}

// ProductUsecase is a Greeter usecase.
type ProductUsecase struct {
	repo ProductRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ProductUsecase) CreateProduct(ctx context.Context, g *Product) (*Product, error) {
	// uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
