package grpc

import (
	"context"

	"github.com/belmadge/goCleanArchitecture/pkg/services"
	pb "github.com/belmadge/goCleanArchitecture/proto"
)

type OrderGRPCService struct {
	pb.UnimplementedOrderServiceServer
	OrderService *services.OrderService
}

func (s *OrderGRPCService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.OrderService.ListOrders()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:           int32(order.ID),
			CustomerName: order.CustomerName,
			OrderDate:    order.OrderDate,
			Status:       order.Status,
		})
	}
	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
