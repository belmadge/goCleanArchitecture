package graphql

import (
	"context"
	"strconv"

	"github.com/belmadge/goCleanArchitecture/pkg/api/graphql/generated"
	"github.com/belmadge/goCleanArchitecture/pkg/api/graphql/model"
	"github.com/belmadge/goCleanArchitecture/pkg/services"
)

type Resolver struct {
	OrderService *services.OrderService
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (q *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	orders, err := q.OrderService.ListOrders()
	if err != nil {
		return nil, err
	}

	var gqlOrders []*model.Order
	for _, order := range orders {
		gqlOrders = append(gqlOrders, &model.Order{
			ID:           strconv.Itoa(order.ID),
			CustomerName: order.CustomerName,
			OrderDate:    order.OrderDate,
			Status:       order.Status,
		})
	}
	return gqlOrders, nil
}
