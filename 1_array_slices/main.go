package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Name string
	Price float64
	Quantity int
}

var products []Product = []Product{
	{Name: "Iphone 12", Price: 24.120, Quantity: 120},
	{Name: "MacBook Pro M3Pro", Price: 120.000, Quantity: 20},
	{Name: "Apple Watch SE", Price: 12.890, Quantity: 1000},
}

func (p Product) toString () string {
	return fmt.Sprintf("Name: %s , Price %.2f, Count: %d\n", p.Name, p.Price, p.Quantity)
}

func main() {
	var search string

	for true {
		fmt.Println("Enter product for search:");

		fmt.Scanf("%s", &search)
	
		foundProduct := getProductByName(search)
	
		if (foundProduct == nil) {
			fmt.Println("Product not found")
		} else {
			fmt.Println(foundProduct.toString())
		}

		fmt.Println("---------")
	}
	
}

func getProductByName(productName string) *Product {
	for _, product := range products {
		if strings.Contains(product.Name, productName) {
			return &product
		}
	}

	return nil
}