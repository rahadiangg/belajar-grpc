package service

import (
	"context"

	"github.com/rahadiangg/belajar-grpc/model"
	"github.com/rahadiangg/belajar-grpc/pkg"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	model.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProducts(ctx context.Context, pageParam *model.Page) (*model.Products, error) {

	var page int64 = 1

	if pageParam.GetPage() != 0 {
		page = pageParam.GetPage()
	}

	var pagination model.Pagination
	var products []*model.Product

	sql := p.DB.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id as category_id", "c.name as category_name")

	offset, limit := pkg.Pagination(sql, page, &pagination)

	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for rows.Next() {
		var product model.Product
		var category model.Category

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
			log.Error("Can't get rows data: ", err.Error())
		}

		product.Category = &category
		products = append(products, &product)
	}

	response := &model.Products{
		Pagination: &pagination,
		Data:       products,
	}

	return response, nil

}
