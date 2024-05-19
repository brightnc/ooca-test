package main

import (
	"fmt"
)

const (
	memberDiscountPercentage      = 10
	orderDoubleDiscountPercentage = 5
)

type Calculator struct {
	menu       map[string]float64
	memberCard bool
	orders     map[string]int
}

func NewCalculator(memberCard bool) *Calculator {
	return &Calculator{
		menu: map[string]float64{
			"Red":    50.0,
			"Green":  40.0,
			"Blue":   30.0,
			"Yellow": 50.0,
			"Pink":   80.0,
			"Purple": 90.0,
			"Orange": 120.0,
		},
		memberCard: memberCard,
		orders:     make(map[string]int),
	}
}

func (c *Calculator) AddOrder(item string, quantity int) {
	if _, exists := c.menu[item]; exists {
		c.orders[item] += quantity
	} else {
		fmt.Printf("Item %v does not exist on the menu", item)
	}
}

func (c *Calculator) CalculateTotal() (float64, int) {
	var totalPrice float64
	totalDiscountPercentage := 0

	for item, quantity := range c.orders {
		totalPrice += c.menu[item] * float64(quantity)
	}

	if c.memberCard {
		totalDiscountPercentage += memberDiscountPercentage
	}
	itemsPromotion := []string{"Orange", "Pink", "Green"}
	for _, item := range itemsPromotion {
		quantity, exists := c.orders[item]
		if exists && quantity > 1 {
			totalDiscountPercentage += orderDoubleDiscountPercentage
			break
		}
	}

	totalPrice = totalPrice * float64(100-totalDiscountPercentage) / 100

	return totalPrice, totalDiscountPercentage
}

func main() {
	calculator := NewCalculator(true)
	calculator.AddOrder("Red", 1)
	calculator.AddOrder("Green", 1)
	total, totalDiscountPercentage := calculator.CalculateTotal()
	fmt.Printf("Order: %v ,Total price: %v, total discount: %v\n", calculator.orders, total, totalDiscountPercentage)

	calculator = NewCalculator(true)
	calculator.AddOrder("Orange", 3)
	calculator.AddOrder("Pink", 4)
	calculator.AddOrder("Green", 4)
	total, totalDiscountPercentage = calculator.CalculateTotal()
	fmt.Printf("Order: %v ,Total price: %v, total discount: %v\n", calculator.orders, total, totalDiscountPercentage)

}
