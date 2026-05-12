package main

import (
	"fmt"
	"strings"
)

var productPrice = map[string]float64{
	"Samsung":     100.0,
	"Iphone":      150.0,
	"Google Pixel": 200.0,
	"OnePlus":     120.0,
	"Sony":        130.0,
	"LG":          110.0,
}

func calculateItemPrice(item string) (float64, bool) {

	basePrice, found := productPrice[item]

	if !found {

		if strings.HasSuffix(item, "_SALE") {

			originalItem := strings.TrimSuffix(item, "_SALE")

			basePrice, found = productPrice[originalItem]

			if found {

				salePrice := basePrice * 0.8 // 20% discount

				fmt.Printf(
					"Item: %s, Original Price: %.2f, Sale Price: %.2f\n",
					item,
					basePrice,
					salePrice,
				)

				return salePrice, true
			}
		}

		fmt.Printf("Item: %s not found in product list.\n", item)

		return 0.0, false
	}

	return basePrice, true
}

func main() {

	fmt.Println("Welcome to the Sales Order Process!")

	orderItems := []string{
		"Samsung",
		"Iphone",
		"Google Pixel",
		"OnePlus",
		"Sony",
		"LG",
		"Unknown Product",
		"Samsung_SALE",
	}

	var totalPrice float64

	fmt.Println("Calculating prices for order items...")

	for _, item := range orderItems {

		price, found := calculateItemPrice(item)

		if found {
			totalPrice += price
		}
	}

	fmt.Printf("Total Price for the order: %.2f\n", totalPrice)
}