package http

import (
    "encoding/json"
    "net/http"
    "github.com/belmadge/goCleanArchitecture/pkg/repository"
)

type OrderHandler struct {
    repo repository.OrderRepository
}

func NewOrderHandler(repo repository.OrderRepository) *OrderHandler {
    return &OrderHandler{repo: repo}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
    orders, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}
