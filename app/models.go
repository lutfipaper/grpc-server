package app

import (
	"context"

	proto "github.com/lutfipaper/module-proto/go/services/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type Models interface {
	GetListProduct(ctx context.Context, req *emptypb.Empty) (res *proto.GetListProductResponse, err error)
}

type models struct {
	db *gorm.DB
}

func NewModels() Models {
	return &models{
		db: GetMysqlConnection(),
	}
}
func (c *models) GetListProduct(ctx context.Context, req *emptypb.Empty) (res *proto.GetListProductResponse, err error) {
	res = &proto.GetListProductResponse{}
	var result []ProductList

	tx := c.db.Raw("SELECT * FROM products").Debug().Model(&ProductList{}).Scan(&result)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return res, status.Error(codes.NotFound, tx.Error.Error())
		}
		return res, status.Error(codes.Internal, tx.Error.Error())
	}

	for _, r := range result {
		res.ListProduct = append(res.ListProduct, &proto.Product{
			UUID:        r.UUID,
			Name:        r.Name,
			Description: r.Description,
			Price:       r.SalesPrice,
		})
	}
	return res, err
}
