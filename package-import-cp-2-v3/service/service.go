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

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

// Implement methods of ServiceInterface

func (s *Service) AddCart(productName string, quantity int) error {
	// Implement logic to add a product to the cart
	return nil
}

func (s *Service) RemoveCart(productName string) error {
	// Implement logic to remove a product from the cart
	return nil
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	// Implement logic to show the cart
	return nil, nil
}

func (s *Service) ResetCart() error {
	// Implement logic to reset the cart
	return nil
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	// Implement logic to get all products
	return nil, nil
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	// Implement logic to process payment
	return entity.PaymentInformation{}, nil
}
