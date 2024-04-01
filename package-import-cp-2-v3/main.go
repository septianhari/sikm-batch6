package main

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/service"
	"fmt"
)

func CashierApp(db *database.Database) service.ServiceInterface {
	service := service.NewService(db)
	return service
}

func main() {
	database := database.NewDatabase()
	service := CashierApp(database)

	// Test Case: Add valid product to cart
	err := service.AddCart("Kaos Polos", 2)
	if err != nil {
		panic(err)
	}

	// Test Case: ShowCart
	fmt.Println("Cart after adding products:")
	cartItems, err := service.ShowCart()
	if err != nil {
		panic(err)
	}
	for _, item := range cartItems {
		fmt.Printf("Product Name: %s, Price: %d, Quantity: %d\n", item.ProductName, item.Price, item.Quantity)
	}

	// Test Case: Pay
	paymentInformation, err := service.Pay(500000)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nPayment Information:")
	fmt.Printf("Total Price: %d\n", paymentInformation.TotalPrice)
	fmt.Printf("Money Paid: %d\n", paymentInformation.MoneyPaid)
	fmt.Printf("Change: %d\n", paymentInformation.Change)
}
