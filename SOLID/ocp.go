package main

//open-close principle ------> open to extension-close to modification
import "fmt"

type Color int

const (
	green Color = iota
	red
	blue
)

type Size int

const (
	large Size = iota
	small
)

type Product struct {
	name  string
	color Color
	size  Size
}

// Filter Bad example
type Filter struct{}

func (f Filter) FilterByColor(p []Product, color Color) []Product {
	var result []Product
	for _, v := range p {
		if v.color == color {
			result = append(result, v)
		}
	}
	return result
}
func (f Filter) FilterBySize(p []Product, size Size) []Product {
	var result []Product
	for i, v := range p {
		if v.size == size {
			result = append(result, p[i])
		}
	}
	return result
}

// Specification --------------------------------------------------------------------------
//Good example (Using specification pattern)
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (colorSpec ColorSpecification) IsSatisfied(p *Product) bool {
	return colorSpec.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (sizeSpec SizeSpecification) IsSatisfied(p *Product) bool {
	return sizeSpec.size == p.size
}

type BetterFilter struct{}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	products := []Product{
		{"Apple", green, large},
		{"Watch", red, large},
		{"House", blue, small},
	}
	f := Filter{}
	fmt.Printf("- filter result(old): %v\n", f.FilterByColor(products, blue))
	fmt.Printf("- filter result(old): %v\n", f.FilterBySize(products, large))

	// vvv AFTER
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}
}
