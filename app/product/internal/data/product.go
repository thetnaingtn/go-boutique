package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/thetnaingtn/go-boutique/app/product/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("product not found")
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

func (r *productRepo) Update(ctx context.Context, id string, p *biz.UpdateProduct) (*biz.Product, error) {
	product, err := r.FindByID(ctx, id)
	if err != nil {
		return &biz.Product{}, err
	}

	if p.Name != nil {
		product.Name = *p.Name
	}

	if p.Description != nil {
		product.Description = *p.Description
	}

	if p.Price != nil {
		product.Price = *p.Price
	}

	q := `UPDATE products SET "name" = :name, "description" = :description, "price" = :price WHERE "id" = :id`

	_, err = r.data.db.NamedExecContext(ctx, q, product)

	if err != nil {
		return &biz.Product{}, err
	}

	return product, nil
}

func (r *productRepo) FindByID(ctx context.Context, id string) (*biz.Product, error) {
	var product biz.Product
	data := struct {
		Id string `db:"id"`
	}{Id: id}
	row, _ := r.data.db.NamedQueryContext(ctx, `SELECT * FROM products WHERE id = :id`, data)

	if !row.Next() {
		return &biz.Product{}, ErrProductNotFound
	}

	if err := row.StructScan(&product); err != nil {
		return &biz.Product{}, fmt.Errorf("error when getting user by id: %w", err)
	}

	return &product, nil
}

func (r *productRepo) DeleteById(ctx context.Context, id string) (string, error) {
	product, err := r.FindByID(ctx, id)

	if err != nil {
		return "", err
	}

	if product == nil {
		return "", ErrProductNotFound
	}

	q := `DELETE FROM products WHERE id = :id`

	data := struct{ Id string }{Id: id}

	_, err = r.data.db.NamedExecContext(ctx, q, data)
	if err != nil {
		return "", err
	}

	return id, nil

}

func (r *productRepo) ListByHello(context.Context, string) ([]*biz.Product, error) {
	return nil, nil
}

func (r *productRepo) ListAll(context.Context) ([]*biz.Product, error) {
	return nil, nil
}
