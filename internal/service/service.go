package service

import (
	v1 "Link-Kratos/api/link/v1"
	"Link-Kratos/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewLinkService)

type LinkService struct {
	v1.UnimplementedLinkServer

	uc *biz.TrafficUsecase
}

func NewLinkService(uc *biz.TrafficUsecase) *LinkService {
	return &LinkService{uc: uc}
}
