package grpc

import (
    "context"
    pb "github.com/belmadge/goCleanArchitecture/proto"
    "github.com/belmadge/goCleanArchitecture/internal/app/service"
)

type OrderServiceServer struct {
    pb.UnimplementedOrderServiceServer
    service *service.OrderService
}

func NewOrderServiceServer(service *service.OrderService) *OrderServiceServer {
    return &OrderServiceServer{service: service}
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
    orders, err := s.service.ListOrders()
    if err != nil {
        return nil, err
    }

    var grpcOrders []*pb.Order
    for _, order := range orders {
        grpcOrders = append(grpcOrders, &pb.Order{
            Id:           int32(order.ID),
            CustomerName: order.CustomerName,
            Total:        order.Total,
        })
    }

    return &pb.ListOrdersResponse{Orders: grpcOrders}, nil
}
