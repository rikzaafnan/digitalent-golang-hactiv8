package itemrepository

import (
	"assignment-2/dto"
	"assignment-2/entity"
)

type ItemRepository interface {
	CreateItem(req *dto.ItemRequest) (int64, error)
	FindItemById(itemId int64) (entity.Item, error)
	FindAllItemByOrderID(orderID int64) ([]entity.Item, error)
	UpdateItem(itemId int64, req *dto.ItemRequest) (int64, error)
	DeleteItem(itemId int64) (bool, error)
}
