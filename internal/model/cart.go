package model

import "github.com/pkg/errors"

type Cart struct {
	Id        int64 `db:"id"`
	CreatedAt int64 `db:"created_at"`
}

var (
	errInvalidCartId = errors.New("wrong product id value, must be greater than 0")
)

func (c *Cart) CheckId() error {
	if c.Id < 1 {
		return errInvalidCartId
	}

	return nil
}
