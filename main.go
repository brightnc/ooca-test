package main

import (
	"fmt"
	"math"
)

const (
	memberDiscountPercentage      = 10
	orderDoubleDiscountPercentage = 5
)

type Calculator struct {
	menu       map[string]int
	memberCard bool
	orders     map[string]int
}

func NewCalculator(memberCard bool) *Calculator {
	return &Calculator{
		menu: map[string]int{
			"Red":    50,
			"Green":  40,
			"Blue":   30,
			"Yellow": 50,
			"Pink":   80,
			"Purple": 90,
			"Orange": 120,
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
	totalPrice := 0.0
	totalDiscountPercentage := 0

	for item, quantity := range c.orders {
		totalPrice += float64(c.menu[item] * quantity)
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

	totalPrice = math.Floor((totalPrice * (100 - float64(totalDiscountPercentage)) / 100))

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
