package memorystore

import (
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

func (s *MemoryProductStore) List() []*model.Product {
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

	return res
}

func (s *MemoryProductStore) Add(p *model.Product) error {
	s.waitChan <- struct{}{}
	s.mux.Lock()
	defer func() {
		s.mux.Unlock()
		<-s.waitChan
	}()

	s.lastId++
	p.SetId(s.lastId)
	if _, ok := s.data[p.GetId()]; ok {
		return errors.Wrapf(ErrProductExist, "id: %d", p.GetId())
	}

	s.data[p.GetId()] = p
	return nil
}

func (s *MemoryProductStore) Delete(id uint) error {
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

func (s MemoryProductStore) Update(p *model.Product) error {
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
