package catalog

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type mockProductRepository struct {
	products []models.Product
}

var _ models.ProductRepositoryInterface = &mockProductRepository{}

func (m *mockProductRepository) GetProducts(offset, limit int, category string, priceLessThan float64) ([]models.Product, int64, error) {
	return nil, 0, nil
}

func (m *mockProductRepository) GetProductByCode(code string) (*models.Product, error) {
	for _, p := range m.products {
		if p.Code == code {
			return &p, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockProductRepository) GetAllProducts() ([]models.Product, error) {
	return m.products, nil
}

func TestCatalogHandler_HandleGetProductByCode(t *testing.T) {
	repo := &mockProductRepository{
		products: []models.Product{
			{
				ID:    1,
				Code:  "PROD001",
				Price: decimal.NewFromFloat(10.0),
				Category: models.Category{
					ID:   1,
					Code: "clothing",
					Name: "Clothing",
				},
			},
		},
	}

	h := NewCatalogHandler(repo)

	t.Run("product found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog/{code}", nil)
		req.SetPathValue("code", "PROD001")
		rr := httptest.NewRecorder()

		h.HandleGetProductByCode(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{"ID":1,"Code":"PROD001","Price":"10","CategoryID":0,"Category":{"ID":1,"Code":"clothing","Name":"Clothing"},"Variants":null}`, rr.Body.String())
	})

	t.Run("product not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/catalog/{code}", nil)
		req.SetPathValue("code", "PROD002")
		rr := httptest.NewRecorder()

		h.HandleGetProductByCode(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.JSONEq(t, `{"error":"product not found"}`, rr.Body.String())
	})
}
