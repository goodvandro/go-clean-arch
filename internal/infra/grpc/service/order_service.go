package service

import (
	"context"

	"github.com/goodvandro/go-clean-arch/internal/infra/grpc/pb"
	"github.com/goodvandro/go-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderService usecase.CreateOrderUseCase
}

func NewOrderService(createOrderService usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderService: createOrderService,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderService.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
