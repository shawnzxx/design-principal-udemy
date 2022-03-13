package main

import "fmt"

/* This file includes of demonstration about Open-Closed Principal
combination of OCP and Specification pattern
*/

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

// Filter this type is empty struct, because we only need it as receiver
type Filter struct {
}

// filterByColor consider this is our first completed method it was well tested and currently used in production
func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// filterBySize our manager come to us and ask for adding new filter by size, how do want to make changes, add one more filter method, you are breaking OCP
func (f *Filter) filterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// filterByColorAndSize after we rolled out a filter features for one month, manager asked to add in one more combined filter again!!
func (f *Filter) filterByColorAndSize(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.color == color && product.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

/*
We change to use specification pattern to maintain open-closed principal
use interface to extend the different filter
*/

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// CombinedSpecification combination of two type of specification, use for combined both size and color filter
type CombinedSpecification struct {
	first, second Specification
}

func (spec CombinedSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

// BetterFilter this type is empty struct, because we only need it as receiver
type BetterFilter struct {
}

func (f *BetterFilter) FilterBySpecification(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if spec.IsSatisfied(&product) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	ball := Product{"Ball", blue, small}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, ball, house}

	// vvv BEFORE
	fmt.Print("Green products (old implementation):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// vvv AFTER
	fmt.Print("Green products (new implementation):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, product := range bf.FilterBySpecification(products, greenSpec) {
		fmt.Printf(" - %s is green\n", product.name)
	}

	fmt.Print("Large products (new implementation):\n")
	largeSpec := SizeSpecification{large}
	for _, product := range bf.FilterBySpecification(products, largeSpec) {
		fmt.Printf(" - %s is large\n", product.name)
	}

	fmt.Print("Large and Green products (new implementation):\n")
	combinedSpec := CombinedSpecification{greenSpec, largeSpec}
	for _, product := range bf.FilterBySpecification(products, combinedSpec) {
		fmt.Printf(" - %s is large and green\n", product.name)
	}
}
