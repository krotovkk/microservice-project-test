package memorystore

import (
	"github.com/pkg/errors"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

var (
	errProductExist    = errors.New("product exist")
	errProductNotExist = errors.New("product doesn't exist")
)

type MemoryProductStore struct {
	data   map[uint]*model.Product
	lastId uint
}

func NewMemoryProductStore() *MemoryProductStore {
	return &MemoryProductStore{data: map[uint]*model.Product{}}
}

func (s *MemoryProductStore) List() []*model.Product {
	res := make([]*model.Product, 0, len(s.data))

	for _, v := range s.data {
		res = append(res, v)
	}

	return res
}

func (s *MemoryProductStore) Add(p *model.Product) error {
	s.lastId++
	p.SetId(s.lastId)
	if _, ok := s.data[p.GetId()]; ok {
		return errors.Wrapf(errProductExist, "id: %d", p.GetId())
	}

	s.data[p.GetId()] = p
	return nil
}

func (s *MemoryProductStore) Delete(id uint) error {
	if _, ok := s.data[id]; !ok {
		return errors.Wrapf(errProductNotExist, "id: %d", id)
	}

	delete(s.data, id)
	return nil
}

func (s MemoryProductStore) Update(p *model.Product) error {
	if _, ok := s.data[p.GetId()]; !ok {
		return errors.Wrapf(errProductNotExist, "id: %d", p.GetId())
	}

	s.data[p.GetId()] = p
	return nil
}
