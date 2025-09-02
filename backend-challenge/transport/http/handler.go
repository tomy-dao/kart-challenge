package httpTransoprt

import (
	"encoding/json"
	"local/endpoint"
	"local/model"
	"local/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	endpoints *endpoint.Endpoints
}

func (h *handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":    "OK",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"service":   "local-service",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	productRes := h.endpoints.Product.GetAllProducts()
	utils.HandlerResponse(w, productRes)
}

func (h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	product := h.endpoints.Product.GetProductById(chi.URLParam(r, "id"))
	utils.HandlerResponse(w, product)
}

func (h *handler) PlaceAnOrder(w http.ResponseWriter, r *http.Request) {
	var req model.OrderRequest
	json.NewDecoder(r.Body).Decode(&req)
	orderRes := h.endpoints.Order.PlaceAnOrder(req)

	if orderRes.Error() {
		w.WriteHeader(orderRes.StatusCode())
		json.NewEncoder(w).Encode(model.ErrorResponse(orderRes))
		return
	} else {
		utils.HandlerResponse(w, orderRes)
	}
}
