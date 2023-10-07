package data

import (
	"github.com/thetnaingtn/go-boutique/app/product/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo)

// Data .
type Data struct {
	db *sqlx.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	var db *sqlx.DB

	cleanup := func() {
		if db != nil {
			db.Close()
		}
		log.NewHelper(logger).Info("closing the data resources")
	}

	db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)

	if err != nil {
		return &Data{}, cleanup, err
	}

	data := &Data{
		db,
	}

	return data, cleanup, nil
}
