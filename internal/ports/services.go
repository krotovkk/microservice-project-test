package ports

import "gitlab.ozon.dev/krotovkk/homework/internal/model"

type ProductService interface {
	List() []*model.Product
	Create(name string, price float64) error
	Update(id uint, name string, price float64) error
	Delete(id uint) error
}
