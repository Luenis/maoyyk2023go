package main

import (
	"fmt"
)

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

type Describable interface {
	Description() string
}

func (item Item) Description() string {
	discountedPrice := item.Price * (1 - item.Discount)
	if item.Discount > 0 {
		return fmt.Sprintf("%s - %.2f (Indirimli fiyat: %.2f)", item.Name, item.Price, discountedPrice)
	}
	return fmt.Sprintf("%s - %.2f", item.Name, item.Price)
}

func calculatePrice(item Item) float64 {
	discountedPrice := item.Price * (1 - item.Discount)
	return discountedPrice
}

func totalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += calculatePrice(item)
	}
	return total
}

func main() {

	items := []Item{
		{"Ekmek", 10.0, 0.1},
		{"Portakal", 20.0, 0.2},
		{"Köfte", 30.0, 0.15},
	}

	fmt.Println("Items:")
	//
	for _, item := range items {
		fmt.Println(item.Description())
	}

	fmt.Printf("\nToplam Fiyat %.2f\n", totalPrice(items))
}
