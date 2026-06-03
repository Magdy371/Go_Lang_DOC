package main

import "fmt"

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func calculateTotal(cart []CartItem) float64 {
	var total float64 = 0
	for _, v := range cart {
		total += v.Price * float64(v.Quantity)
	}
	return total
}
func main() {
	var cart []CartItem
	var apples = CartItem{"apple", 2.99, 3}
	var oranges = CartItem{"orange", 1.50, 8}
	cart = append(cart, apples, oranges)
	total := calculateTotal(cart)
	fmt.Printf("total = %f\n", total)
}
