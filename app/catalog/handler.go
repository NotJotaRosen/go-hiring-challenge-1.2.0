package catalog

import (
	"net/http"
	"strconv"

	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/models"
	"gorm.io/gorm"
)

type response struct {
	Products []models.Product `json:"products"`
	Total    int64            `json:"total"`
}



type CatalogHandler struct {
	repo models.ProductRepositoryInterface
}

func NewCatalogHandler(r models.ProductRepositoryInterface) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	category := r.URL.Query().Get("category")
	priceLessThan, _ := strconv.ParseFloat(r.URL.Query().Get("priceLessThan"), 64)

	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	res, total, err := h.repo.GetProducts(offset, limit, category, priceLessThan)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, response{Products: res, Total: total})
}

func (h *CatalogHandler) HandleGetProductByCode(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	product, err := h.repo.GetProductByCode(code)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			api.ErrorResponse(w, http.StatusNotFound, "product not found")
			return
		}
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, product)
}
