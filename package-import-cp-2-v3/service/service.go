package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
)

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

// GetAllProduct implements ServiceInterface.
func (s *Service) GetAllProduct() ([]entity.Product, error) {
	panic("unimplemented")
}

// Pay implements ServiceInterface.
func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	panic("unimplemented")
}

// RemoveCart implements ServiceInterface.
func (s *Service) RemoveCart(productName string) error {
	panic("unimplemented")
}

// ResetCart implements ServiceInterface.
func (s *Service) ResetCart() error {
	panic("unimplemented")
}

// ShowCart implements ServiceInterface.
func (s *Service) ShowCart() ([]entity.CartItem, error) {
	panic("unimplemented")
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	// Retrieve existing cart items
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	// Retrieve product information
	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}

	// Check if the product already exists in the cart
	for i, item := range cartItems {
		if item.ProductName == productName {
			cartItems[i].Quantity += quantity
			// Update the cart in the database
			err := s.database.SaveCartItems(cartItems)
			if err != nil {
				return err
			}
			return nil
		}
	}

	// If the product is not in the cart, add it
	newItem := entity.CartItem{
		ProductName: productName,
		Price:       product.Price,
		Quantity:    quantity,
	}
	cartItems = append(cartItems, newItem)

	// Update the cart in the database
	err = s.database.SaveCartItems(cartItems)
	if err != nil {
		return err
	}

	return nil
}

// Implement the other methods similarly
