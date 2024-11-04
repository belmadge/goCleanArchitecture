package services

import (
	"database/sql"

	"github.com/belmadge/goCleanArchitecture/pkg/models"
)

type OrderService struct {
	DB *sql.DB
}

func (s *OrderService) ListOrders() ([]models.Order, error) {
	rows, err := s.DB.Query("SELECT id, customer_name, order_date, status FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.CustomerName, &order.OrderDate, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
