package database

import (
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

var productData = []entity.Product{
	{Name: "Kaos Polos", Price: 50000},
	{Name: "Kaos sablon", Price: 70000},
	{Name: "Baju Batik", Price: 200000},
	{Name: "Celana hitam", Price: 150000},
	{Name: "Celana pendek", Price: 100000},
	{Name: "Sabuk", Price: 50000},
	{Name: "Topi", Price: 75000},
	{Name: "Galung Tangan", Price: 30000},
	{Name: "Sepatu", Price: 300000},
}

type DatabaseInterface interface {
	GetCartItems() ([]entity.CartItem, error)
	SaveCartItems([]entity.CartItem) error
	GetProductData() []entity.Product
	GetProductByName(name string) (entity.Product, error)
	RemoveCartItem(productName string) error
	ResetCart() error
}

type Database struct {
	Cart []entity.CartItem
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) GetCartItems() ([]entity.CartItem, error) {
	return db.Cart, nil
}

func (db *Database) SaveCartItems(cartItems []entity.CartItem) error {
	db.Cart = cartItems
	return nil
}

func (db *Database) GetProductData() []entity.Product {
	return productData
}

func (db *Database) GetProductByName(name string) (entity.Product, error) {
	var product entity.Product
	for _, p := range productData {
		if p.Name == name {
			return p, nil
		}
	}

	return product, errors.New("product not found")
}

func (db *Database) RemoveCartItem(productName string) error {
	for i, item := range db.Cart {
		if item.ProductName == productName {
			db.Cart = append(db.Cart[:i], db.Cart[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found in cart")
}

func (db *Database) ResetCart() error {
	db.Cart = nil
	return nil
}
