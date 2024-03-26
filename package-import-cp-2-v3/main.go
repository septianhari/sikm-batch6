package main

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"fmt"
	"log"
)

func main() {
	// Initialize database
	db := database.NewDatabase()

	// Initialize service with database dependency injection
	cashierService := service.NewCashierService(db)

	// Add products to the database
	products := []entity.Product{
		{Name: "Kaos Polos", Price: 50000},
		{Name: "Kaos Sablon", Price: 70000},
		{Name: "Baju Batik", Price: 200000},
		{Name: "Celana Hitam", Price: 150000},
		{Name: "Celana Pendek", Price: 100000},
		{Name: "Sabuk", Price: 50000},
		{Name: "Topi", Price: 75000},
		{Name: "Gelang Tangan", Price: 30000},
		{Name: "Sepatu", Price: 300000},
	}
	for _, product := range products {
		err := db.SaveProductData(product)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Test AddCart function
	err := cashierService.AddCart("Kaos Polos", 2)
	if err != nil {
		log.Fatal(err)
	}

	// Test ShowCart function
	cartItems, err := cashierService.ShowCart()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cart Items:")
	for _, item := range cartItems {
		fmt.Printf("Product: %s, Price: %d, Quantity: %d\n", item.ProductName, item.Price, item.Quantity)
	}

	// Test RemoveCart function
	err = cashierService.RemoveCart("Kaos Polos")
	if err != nil {
		log.Fatal(err)
	}

	// Test ResetCart function
	err = cashierService.ResetCart()
	if err != nil {
		log.Fatal(err)
	}

	// Test GetAllProduct function
	allProducts, err := cashierService.GetAllProduct()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Products:")
	for _, product := range allProducts {
		fmt.Printf("Product: %s, Price: %d\n", product.Name, product.Price)
	}

	// Test Pay function
	err = cashierService.AddCart("Kaos Polos", 2)
	if err != nil {
		log.Fatal(err)
	}
	err = cashierService.AddCart("Kaos Sablon", 1)
	if err != nil {
		log.Fatal(err)
	}
	paymentInfo, err := cashierService.Pay(500000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Payment Information:")
	fmt.Printf("Total Price: %d\n", paymentInfo.TotalPrice)
	fmt.Printf("Money Paid: %d\n", paymentInfo.MoneyPaid)
	fmt.Printf("Change: %d\n", paymentInfo.Change)
	fmt.Println("Product List:")
	for _, item := range paymentInfo.ProductList {
		fmt.Printf("Product: %s, Price: %d, Quantity: %d\n", item.ProductName, item.Price, item.Quantity)
	}
}
