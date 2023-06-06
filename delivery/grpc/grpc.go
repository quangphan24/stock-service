package grpc

import (
	"context"
	"encoding/json"
	proto "stock-service/proto/stock"
	"stock-service/usecase"
)

type ServerGRPC struct {
	proto.UnimplementedStockServiceServer
	UseCase *usecase.UseCase
}

func (s *ServerGRPC) GetProductById(ctx context.Context, in *proto.Int) (*proto.ProductReply, error) {
	product, err := s.UseCase.ProductUseCase.GetProductById(int(in.GetValue()))
	if err != nil {
		return nil, err
	}
	var reply *proto.ProductReply
	data, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
