package service

import (
	"context"

	v1 "Link-Kratos/api/link/v1"
	"Link-Kratos/internal/biz"
)

type LinkService struct {
	v1.UnimplementedLinkServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *LinkService {
	return &LinkService{uc: uc}
}

func (s *LinkService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
