package models

type ProductRepositoryInterface interface {
	GetAllProducts() ([]Product, error)
	GetProducts(offset, limit int, category string, priceLessThan float64) ([]Product, int64, error)
	GetProductByCode(code string) (*Product, error)
}
