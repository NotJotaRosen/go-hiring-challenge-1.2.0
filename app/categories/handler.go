package categories

import (
	"encoding/json"
	"net/http"

	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/models"
)

type CategoriesHandler struct {
	repo models.CategoryRepository
}

func NewCategoriesHandler(r models.CategoryRepository) *CategoriesHandler {
	return &CategoriesHandler{r}
}

func (h *CategoriesHandler) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetAllCategories()
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, categories)
}

func (h *CategoriesHandler) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.repo.CreateCategory(&category); err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, category)
}
