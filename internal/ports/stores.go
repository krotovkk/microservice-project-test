package ports

import "gitlab.ozon.dev/krotovkk/homework/internal/model"

type ProductStore interface {
	List() []*model.Product
	Add(p *model.Product) error
	Delete(id uint) error
	Update(p *model.Product) error
}
