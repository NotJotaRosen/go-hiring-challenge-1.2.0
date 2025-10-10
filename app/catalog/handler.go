package catalog

import (
	"net/http"
	"strconv"

	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/models"
)

type response struct {
	Products []product `json:"products"`
	Total    int64     `json:"total"`
}

type product struct {
	Code     string  `json:"code"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type CatalogHandler struct {
	repo *models.ProductsRepository
}

func NewCatalogHandler(r *models.ProductsRepository) *CatalogHandler {
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

	products := make([]product, len(res))
	for i, p := range res {
		products[i] = product{
			Code:     p.Code,
			Price:    p.Price.InexactFloat64(),
			Category: p.Category.Name,
		}
	}

	api.OKResponse(w, response{Products: products, Total: total})
}
