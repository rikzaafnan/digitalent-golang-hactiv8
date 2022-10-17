package orderrepository

import (
	"assignment-2/dto"
	"assignment-2/entity"
)

type OrderRepository interface {
	CreateOrder(req *dto.OrderRequest) (int64, error)
	FindOrderById(orderId int64) (entity.Order, error)
	FindAllOrder() ([]entity.Order, error)
	UpdateOrder(orderId int64, req *dto.OrderRequestUpdate) (int64, error)
	DeleteOrder(orderId int64) (bool, error)
}
