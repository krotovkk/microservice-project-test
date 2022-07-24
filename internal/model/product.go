package model

import "fmt"

type Product struct {
	id    uint
	price float64
	name  string
}

func NewProduct(id uint, name string, price float64) (*Product, error) {
	p := &Product{}
	if err := p.SetName(name); err != nil {
		return nil, err
	}
	if err := p.SetPrice(price); err != nil {
		return nil, err
	}
	p.SetId(id)

	return p, nil
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) SetName(name string) error {
	if len(name) < 1 || len(name) > 25 {
		return fmt.Errorf("bad name: <%v>", name)
	}
	p.name = name

	return nil
}

func (p *Product) GetId() uint {
	return p.id
}

func (p *Product) SetId(id uint) {
	p.id = id
}

func (p *Product) SetPrice(price float64) error {
	if price <= 0 {
		return fmt.Errorf("bad price: <%v>", price)
	}
	p.price = price

	return nil
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) String() string {
	return fmt.Sprintf("id: <%v>; name: <%v>; price: %.2f\n", p.id, p.name, p.price)
}
