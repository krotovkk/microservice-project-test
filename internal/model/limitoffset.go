package model

type LimitOffset struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}
