package main

type bill struct {
	name  string
	items map[string]float64
	total float64
}

func NewBill(name string) *bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		total: 0,
	}

	return &b
}

func (b bill) computeTotal() {
	b.total = 0
	for _, value := range b.items {
		b.total += value
	}
}

func (b bill) addItem(n string, v float64) {
	b.items[n] = v
}
