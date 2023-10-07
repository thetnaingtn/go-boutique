package data

import (
	"context"

	"github.com/thetnaingtn/go-boutique/app/product/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

// NewProductRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *productRepo) Save(ctx context.Context, g *biz.Product) (*biz.Product, error) {

	id := uuid.New()

	product := &biz.Product{
		Id:          id.String(),
		Name:        g.Name,
		Description: g.Description,
		Price:       g.Price,
	}

	tx := r.data.db.MustBegin()

	if _, err := tx.NamedExec(`INSERT INTO products (id,name,description,price) VALUES (:id,:name,:description,:price)`, product); err != nil {
		return &biz.Product{}, err
	}

	tx.Commit()

	return product, nil
}

func (r *productRepo) Update(ctx context.Context, g *biz.Product) (*biz.Product, error) {
	return g, nil
}

func (r *productRepo) FindByID(context.Context, int64) (*biz.Product, error) {
	return nil, nil
}

func (r *productRepo) ListByHello(context.Context, string) ([]*biz.Product, error) {
	return nil, nil
}

func (r *productRepo) ListAll(context.Context) ([]*biz.Product, error) {
	return nil, nil
}
