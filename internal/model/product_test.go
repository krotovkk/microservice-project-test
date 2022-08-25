package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckId(t *testing.T) {
	t.Run("Valid id", func(t *testing.T) {
		product := &Product{Id: 20}
		err := product.CheckId()
		assert.NoError(t, err)
	})
}

func TestCheckPrice(t *testing.T) {
	t.Run("Valid price", func(t *testing.T) {
		product := &Product{Id: 1, Price: 220}
		err := product.CheckPrice()
		assert.NoError(t, err)
	})
	t.Run("Negative price", func(t *testing.T) {
		product := &Product{Id: 1, Price: -123}
		err := product.CheckPrice()
		assert.Error(t, err)
	})
}

func TestCheckName(t *testing.T) {
	t.Run("Valid product name", func(t *testing.T) {
		product := &Product{Id: 1, Name: "valid"}
		err := product.CheckName()
		assert.NoError(t, err)
	})
	t.Run("Very long name", func(t *testing.T) {
		product := &Product{Id: 1, Name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
		err := product.CheckName()
		assert.Error(t, err)
	})
}
