package repository

import (
    "database/sql"
    "github.com/belmadge/goCleanArchitecture/pkg/models"
)

type OrderRepository interface {
    GetAll() ([]models.Order, error)
}

type orderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
    return &orderRepository{db: db}
}

func (r *orderRepository) GetAll() ([]models.Order, error) {
    rows, err := r.db.Query("SELECT id, customer_name, total FROM orders")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.ID, &order.CustomerName, &order.Total); err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }

    return orders, nil
}
