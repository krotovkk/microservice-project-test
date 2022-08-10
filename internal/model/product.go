package model

import "fmt"

type Product struct {
	Id    uint    `db:"id"`
	Price float64 `db:"price"`
	Name  string  `db:"name"`
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
	return p.Name
}

func (p *Product) SetName(name string) error {
	if len(name) < 1 || len(name) > 25 {
		return fmt.Errorf("bad Name: <%v>", name)
	}
	p.Name = name

	return nil
}

func (p *Product) GetId() uint {
	return p.Id
}

func (p *Product) SetId(id uint) {
	p.Id = id
}

func (p *Product) SetPrice(price float64) error {
	if price <= 0 {
		return fmt.Errorf("bad Price: <%v>", price)
	}
	p.Price = price

	return nil
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) String() string {
	return fmt.Sprintf("Id: <%v>; Name: <%v>; Price: %.2f\n", p.Id, p.Name, p.Price)
}
