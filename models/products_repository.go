package models

import (
	"gorm.io/gorm"
)

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

func (r *ProductsRepository) GetAllProducts() ([]Product, error) {
	var products []Product
	if err := r.db.Preload("Variants").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductsRepository) GetProducts(offset, limit int, category string, priceLessThan float64) ([]Product, int64, error) {
	var products []Product
	var total int64

	tx := r.db.Model(&Product{}).Preload("Variants").Preload("Category")

	if category != "" {
		tx = tx.Joins("JOIN categories ON categories.id = products.category_id AND categories.code = ?", category)
	}

	if priceLessThan > 0 {
		tx = tx.Where("price < ?", priceLessThan)
	}

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := tx.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *ProductsRepository) GetProductByCode(code string) (*Product, error) {
	var product Product
	if err := r.db.Preload("Variants").Preload("Category").First(&product, "code = ?", code).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
