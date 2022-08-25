package memorystore

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

const maxConcurrentOperations = 10

var (
	ErrProductExist    = errors.New("product exist")
	ErrProductNotExist = errors.New("product doesn't exist")
)

type MemoryProductStore struct {
	data     map[uint]*model.Product
	mux      sync.RWMutex
	lastId   uint
	waitChan chan struct{}
}

func NewMemoryProductStore() *MemoryProductStore {
	return &MemoryProductStore{
		data:     map[uint]*model.Product{},
		mux:      sync.RWMutex{},
		waitChan: make(chan struct{}, maxConcurrentOperations),
	}
}

func (s *MemoryProductStore) GetAllProducts(context.Context, uint64, uint64) ([]*model.Product, error) {
	s.waitChan <- struct{}{}
	s.mux.RLock()
	defer func() {
		s.mux.RUnlock()
		<-s.waitChan
	}()

	res := make([]*model.Product, 0, len(s.data))

	for _, v := range s.data {
		res = append(res, v)
	}

	return res, nil
}

func (s *MemoryProductStore) CreateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	s.waitChan <- struct{}{}
	s.mux.Lock()
	defer func() {
		s.mux.Unlock()
		<-s.waitChan
	}()

	s.lastId++
	p.SetId(s.lastId)
	if _, ok := s.data[p.GetId()]; ok {
		return nil, errors.Wrapf(ErrProductExist, "id: %d", p.GetId())
	}

	s.data[p.GetId()] = p
	return nil, nil
}

func (s *MemoryProductStore) DeleteProduct(ctx context.Context, id uint) error {
	s.waitChan <- struct{}{}
	s.mux.Lock()
	defer func() {
		s.mux.Unlock()
		<-s.waitChan
	}()

	if _, ok := s.data[id]; !ok {
		return errors.Wrapf(ErrProductNotExist, "id: %d", id)
	}

	delete(s.data, id)
	return nil
}

func (s *MemoryProductStore) UpdateProduct(ctx context.Context, p *model.Product) error {
	s.waitChan <- struct{}{}
	s.mux.Lock()
	defer func() {
		s.mux.Unlock()
		<-s.waitChan
	}()

	if _, ok := s.data[p.GetId()]; !ok {
		return errors.Wrapf(ErrProductNotExist, "id: %d", p.GetId())
	}

	s.data[p.GetId()] = p
	return nil
}

func (s *MemoryProductStore) GetProductOne(ctx context.Context, id int64) (*model.Product, error) {
	s.waitChan <- struct{}{}
	s.mux.RLock()
	defer func() {
		s.mux.RUnlock()
		<-s.waitChan
	}()

	if product, ok := s.data[uint(id)]; ok {
		return product, nil
	}

	return nil, errors.Wrapf(ErrProductNotExist, "id: %d", id)
}
