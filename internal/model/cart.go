package model

type Cart struct {
	Id        int64 `db:"id"`
	CreatedAt int64 `db:"created_at"`
}
