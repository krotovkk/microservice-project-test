package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Run("Valid cart model", func(t *testing.T) {
		cart := &Cart{Id: 20}
		err := cart.CheckId()
		assert.NoError(t, err)
	})
}
