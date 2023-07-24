package app

import (
	"context"

	proto "github.com/lutfipaper/module-proto/go/services/product"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controllers struct {
	m Models
	proto.UnimplementedServicesServer
}

func NewControllers(m Models) *Controllers {
	return &Controllers{m: m}
}

func (c *Controllers) GetServiceDesc() []*grpc.ServiceDesc {
	return []*grpc.ServiceDesc{&proto.Services_ServiceDesc}
}

func (c *Controllers) GetListProduct(ctx context.Context, req *emptypb.Empty) (res *proto.GetListProductResponse, err error) {
	return c.m.GetListProduct(ctx, req)
}
