package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	// "fmt"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	c.filebasedDb.StoreCategory(*Category)
	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	// return nil // TODO: replace this
	return c.filebasedDb.UpdateCategory(category.ID, *&category)
}

func (c *categoryRepository) Delete(id int) error {
	// return fmt.Errorf("not implement") // TODO: replace this
	err := c.filebasedDb.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	category, err := c.filebasedDb.GetCategoryByID(id)

	if err != nil {
		return nil, err
	}

	return category, err
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	// return nil, nil // TODO: replace this
	category, err := c.filebasedDb.GetCategories()
	if err != nil {
		return nil, err
	}
	return category, nil
}
