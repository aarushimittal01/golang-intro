package interfaces

import (
	"fmt"
	"sort"
)

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

type Products []Product

type ProductComparer func(p1 Product, p2 Product) bool

type ProductComparers map[string]ProductComparer

var productComparers ProductComparers = ProductComparers{
	"id": func(p1 Product, p2 Product) bool {
		return p1.id < p2.id
	},
	"name": func(p1 Product, p2 Product) bool {
		return p1.name < p2.name
	},
	"cost": func(p1 Product, p2 Product) bool {
		return p1.cost < p2.cost
	},
	"units": func(p1 Product, p2 Product) bool {
		return p1.units < p2.units
	},
	"category": func(p1 Product, p2 Product) bool {
		return p1.category < p2.category
	},
}

var currentProductComparer ProductComparer = productComparers["id"]

/* implementation of "Interface" interface in sort package */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return currentProductComparer(products[i], products[j])
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Sort(attrName string) {
	currentProductComparer = productComparers[attrName]
	sort.Sort(products)
}

func Sort_products() {
	products := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationery"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationery"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utensil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationery"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utensil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationery"},
	}
	fmt.Printf("Products\n%v", products)
	products.Sort("id")
	fmt.Printf("\nAfter sorting by id:\n%v", products)
	products.Sort("name")
	fmt.Printf("\nAfter sorting by name:\n%v", products)
	products.Sort("cost")
	fmt.Printf("\nAfter sorting by cost:\n%v", products)
	products.Sort("units")
	fmt.Printf("\nAfter sorting by units:\n%v", products)
	products.Sort("category")
	fmt.Printf("\nAfter sorting by category:\n%v", products)
}
