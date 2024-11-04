package rest

import (
	"encoding/json"
	"net/http"

	"github.com/belmadge/goCleanArchitecture/pkg/services"
)

func ListOrdersHandler(svc *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := svc.ListOrders()
		if err != nil {
			http.Error(w, "Erro ao listar ordens", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}
