package models

import "gorm.io/gorm"

type CategoryRepository interface {
	GetAllCategories() ([]Category, error)
	CreateCategory(category *Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAllCategories() ([]Category, error) {
	var categories []Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) CreateCategory(category *Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}
