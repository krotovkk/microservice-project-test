//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

func TestCreateCart_Db(t *testing.T) {
	errFailedModel := errors.New("CartPostgresStore.CreateCart: insert: ERROR: new row for relation \"carts\" violates check constraint \"price_chk\" (SQLSTATE 23514)")

	tests := []struct {
		name string
		args model.Cart
		err  error
	}{
		{
			name: "create model 1",
			args: model.Cart{CreatedAt: time.Now().UTC().Unix()},
			err:  nil,
		},
		{
			name: "create model 2",
			args: model.Cart{CreatedAt: time.Now().UTC().Unix()},
			err:  nil,
		},
		{
			name: "create model 3",
			args: model.Cart{CreatedAt: time.Now().UTC().Unix()},
			err:  nil,
		},
		{
			name: "create model 4",
			args: model.Cart{},
			err:  errFailedModel,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//arrange
			Db.SetUp(t)
			defer Db.TearDown()
			//act
			cart, err := Db.Store.Cart().CreateCart(context.Background(), &test.args)
			//assert
			if test.err != nil {
				assert.Equal(t, err.Error(), test.err.Error())
			} else {
				require.Nil(t, err)
				assert.Equal(t, test.args, *cart)
			}
		})
	}
}

func TestGetCartProductsMultiple_Db(t *testing.T) {
	t.Run("Get Cart Product Valid", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		insertProducts := []*model.Product{
			{Price: 100, Name: "test1"},
			{Price: 110, Name: "test2"},
			{Price: 120, Name: "test3"},
			{Price: 130, Name: "test4"},
		}

		c, err := Db.Store.Cart().CreateCart(context.Background(), &model.Cart{CreatedAt: time.Now().UTC().Unix()})
		require.NoError(t, err)

		for _, product := range insertProducts {
			p, err := Db.Store.Product().CreateProduct(context.Background(), product)
			require.NoError(t, err)

			err = Db.Store.Cart().AddProductToCart(context.Background(), int64(p.Id), c.Id)
			require.NoError(t, err)
		}

		products, err := Db.Store.Cart().GetCartProducts(context.Background(), c.Id)

		require.NoError(t, err)
		require.Equal(t, insertProducts, products)
	})
}

func TestAddProductToCart_Db(t *testing.T) {
	t.Run("Add Product To Cart Valid", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		product := &model.Product{Price: 210, Name: "test"}
		p, err := Db.Store.Product().CreateProduct(context.Background(), product)
		require.NoError(t, err)

		c, err := Db.Store.Cart().CreateCart(context.Background(), &model.Cart{CreatedAt: time.Now().UTC().Unix()})
		require.NoError(t, err)

		err = Db.Store.Cart().AddProductToCart(context.Background(), int64(p.Id), c.Id)
		require.NoError(t, err)

		err = Db.Store.Cart().AddProductToCart(context.Background(), int64(p.Id), c.Id)
		require.Error(t, err)
	})
}
