package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOrder(t *testing.T) {
	t.Run("Case : Add item to order", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(false)
		calculator.AddOrder("Red", 1)
		expectedQty := 1
		assert.Equal(1, calculator.orders["Red"], "Expected item quantity %v", expectedQty)
	})

	t.Run("Case : Add multiple item to order", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(false)
		calculator.AddOrder("Red", 2)
		calculator.AddOrder("Green", 1)
		calculator.AddOrder("Blue", 1)
		calculator.AddOrder("Red", 3)
		expectedGreenQty := 1
		expectedRedQty := 5
		expectedBlueQty := 1
		assert.Equal(1, calculator.orders["Green"], "Expected item quantity %v", expectedGreenQty)
		assert.Equal(1, calculator.orders["Blue"], "Expected item quantity %v", expectedBlueQty)
		assert.Equal(5, calculator.orders["Red"], "Expected item quantity %v", expectedRedQty)
	})

	t.Run("Case : Add unexpected menu item", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(false)
		calculator.AddOrder("Coke", 2)
		assert.NotContains(calculator.orders, "Coke", "Expected no Coke in orders")
	})

}
func TestCalculateTotal(t *testing.T) {
	t.Run("Case : With out member card, no promotions", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(false)
		calculator.AddOrder("Red", 1)
		calculator.AddOrder("Green", 1)
		total, totalDiscountPercentage := calculator.CalculateTotal()
		expectedTotal := 50.0 + 40.0
		expectedDiscount := 0
		assert.Equal(expectedTotal, total, "Expected total %v", expectedTotal)
		assert.Equal(expectedDiscount, totalDiscountPercentage, "Expected discount %v", expectedDiscount)
	})

	t.Run("Case : With member card, no promotions", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(true)
		calculator.AddOrder("Blue", 2)
		calculator.AddOrder("Purple", 1)
		total, totalDiscountPercentage := calculator.CalculateTotal()
		expectedTotal := (60.0 + 90.0) * 90 / 100 // 10% discount
		expectedDiscount := 10
		assert.Equal(expectedTotal, total, "Expected total %v", expectedTotal)
		assert.Equal(expectedDiscount, totalDiscountPercentage, "Expected discount %v", expectedDiscount)
	})

	t.Run("Case : With out member card, have promotions", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(false)
		calculator.AddOrder("Orange", 3)
		calculator.AddOrder("Yellow", 1)
		total, totalDiscountPercentage := calculator.CalculateTotal()
		expectedTotal := (360.0 + 50.0) * 95 / 100 // 5% discount
		expectedDiscount := 5
		assert.Equal(expectedTotal, total, "Expected total %v", expectedTotal)
		assert.Equal(expectedDiscount, totalDiscountPercentage, "Expected discount %v", expectedDiscount)
	})

	t.Run("Case : With member card, have promotions", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(true)
		calculator.AddOrder("Orange", 3)
		calculator.AddOrder("Yellow", 1)
		total, totalDiscountPercentage := calculator.CalculateTotal()
		expectedTotal := (360.0 + 50.0) * 85 / 100 // 15% discount
		t.Logf("expectedTotal : %f", expectedTotal)
		expectedDiscount := 15
		assert.Equal(expectedTotal, total, "Expected total %v", expectedTotal)
		assert.Equal(expectedDiscount, totalDiscountPercentage, "Expected discount %v", expectedDiscount)
	})

	t.Run("Case : With member card, multiple promotions items", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(true)
		calculator.AddOrder("Pink", 4)
		calculator.AddOrder("Purple", 2)
		total, totalDiscountPercentage := calculator.CalculateTotal()
		expectedTotal := (320.0 + 180.0) * 85 / 100 // 15% discount
		expectedDiscount := 15
		assert.Equal(expectedTotal, total, "Expected total %v", expectedTotal)
		assert.Equal(expectedDiscount, totalDiscountPercentage, "Expected discount %v", expectedDiscount)
	})

	t.Run("Case : With out member card, multiple promotions items", func(t *testing.T) {
		assert := assert.New(t)
		calculator := NewCalculator(false)
		calculator.AddOrder("Pink", 4)
		calculator.AddOrder("Purple", 2)
		total, totalDiscountPercentage := calculator.CalculateTotal()
		expectedTotal := (320.0 + 180.0) * 95 / 100 // 5% discount
		expectedDiscount := 5
		assert.Equal(expectedTotal, total, "Expected total %v", expectedTotal)
		assert.Equal(expectedDiscount, totalDiscountPercentage, "Expected discount %v", expectedDiscount)
	})

}
