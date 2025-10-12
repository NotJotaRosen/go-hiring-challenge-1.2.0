package categories

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/stretchr/testify/assert"
)

type mockCategoryRepository struct {
	categories []models.Category
}

func (m *mockCategoryRepository) GetAllCategories() ([]models.Category, error) {
	return m.categories, nil
}

func (m *mockCategoryRepository) CreateCategory(category *models.Category) error {
	m.categories = append(m.categories, *category)
	return nil
}

var _ models.CategoryRepository = &mockCategoryRepository{}

func TestCategoriesHandler_HandleGetCategories(t *testing.T) {
	repo := &mockCategoryRepository{
		categories: []models.Category{
			{ID: 1, Code: "clothing", Name: "Clothing"},
			{ID: 2, Code: "shoes", Name: "Shoes"},
		},
	}

	h := NewCategoriesHandler(repo)

	req := httptest.NewRequest("GET", "/categories", nil)
	rr := httptest.NewRecorder()

	h.HandleGetCategories(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `[{"ID":1,"Code":"clothing","Name":"Clothing"},{"ID":2,"Code":"shoes","Name":"Shoes"}]`, rr.Body.String())
}

func TestCategoriesHandler_HandleCreateCategory(t *testing.T) {
	repo := &mockCategoryRepository{}

	h := NewCategoriesHandler(repo)

	category := models.Category{Code: "accessories", Name: "Accessories"}
	body, _ := json.Marshal(category)

	req := httptest.NewRequest("POST", "/categories", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	h.HandleCreateCategory(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"ID":0,"Code":"accessories","Name":"Accessories"}`, rr.Body.String())
}
