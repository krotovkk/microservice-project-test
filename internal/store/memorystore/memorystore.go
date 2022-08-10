package memorystore

import (
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type MemoryStore struct {
	productStore ports.ProductStore
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		productStore: NewMemoryProductStore(),
	}
}

func (s *MemoryStore) Product() ports.ProductStore {
	return s.productStore
}
