package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"log"

	"go.etcd.io/bbolt"
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
	categoryJSON, err := json.Marshal(category)
	if err != nil {
		return err
	}

	return c.filebasedDb.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Categories"))
		return b.Put([]byte(fmt.Sprintf("%d", id)), categoryJSON)
	})
}

func (c *categoryRepository) Delete(id int) error {
	return c.filebasedDb.DeleteCategory(id)
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	category, err := c.filebasedDb.GetCategoryByID(id)

	return category, err
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	var categories []model.Category
	err := c.filebasedDb.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Categories"))
		if b == nil {
			return fmt.Errorf("categories bucket not found")
		}
		return b.ForEach(func(k, v []byte) error {
			var category model.Category
			if err := json.Unmarshal(v, &category); err != nil {
				log.Printf("Error unmarshaling category: %v", err)
				return nil // Continue processing next item in case of error
			}
			categories = append(categories, category)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return categories, nil
}
