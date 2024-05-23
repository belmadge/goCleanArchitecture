package service

import (
    "github.com/belmadge/goCleanArchitecture/pkg/models"
    "github.com/belmadge/goCleanArchitecture/pkg/repository"
)

type OrderService struct {
    repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
    return &OrderService{repo: repo}
}

func (s *OrderService) ListOrders() ([]models.Order, error) {
    return s.repo.GetAll()
}
