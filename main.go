package main

import (
	"fmt"
	pdt "github.com/sunao-uehara/go-github-actions/products"
)

// main
func main() {
	product, err := pdt.GetProduct("p1")
	if err != nil {
		panic("product not found")
	}

	fmt.Printf("%v\n", product)
	product.Name()
}
