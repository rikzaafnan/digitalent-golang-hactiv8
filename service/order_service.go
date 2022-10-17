package service

import (
	"assignment-2/dto"
	itemrepository "assignment-2/repository/item_repository"
	"assignment-2/repository/orderrepository"
	"errors"
	"fmt"
)

type OrderService interface {
	Create(req *dto.OrderRequest) (dto.OrderResponse, error)
	Read(orderID int64) (dto.OrderResponse, error)
	Update(orderID int64, req *dto.OrderRequestUpdate) (dto.OrderResponse, error)
	Delete(orderID int64) error
	GetAll() ([]dto.OrderResponse, error)
}

type orderService struct {
	orderRepository orderrepository.OrderRepository
	itemRepository  itemrepository.ItemRepository
}

func NewOrderService(orderRepository orderrepository.OrderRepository, itemRepository itemrepository.ItemRepository) *orderService {
	return &orderService{
		orderRepository: orderRepository,
		itemRepository:  itemRepository,
	}
}

func (s *orderService) Create(req *dto.OrderRequest) (dto.OrderResponse, error) {

	var orderResponse dto.OrderResponse

	lastInsertId, err := s.orderRepository.CreateOrder(req)
	if err != nil {
		fmt.Println(err)
		return orderResponse, err
	}

	// entityOrder, err := s.orderRepository.FindOrderById(lastInsertId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return orderResponse, err
	// }

	// orderResponse.ID = entityOrder.ID
	// orderResponse.OrderedAt = entityOrder.OrderedAt.GoString()

	orderResponse, err = s.Read(lastInsertId)
	if err != nil {
		fmt.Println(err)
		return orderResponse, err
	}

	return orderResponse, nil

}

func (s *orderService) Read(orderID int64) (dto.OrderResponse, error) {

	var orderResponse dto.OrderResponse

	entityOrder, err := s.orderRepository.FindOrderById(orderID)
	if err != nil {
		fmt.Println(err)
		return orderResponse, err
	}

	orderResponse.ID = entityOrder.ID
	orderResponse.CustomerName = entityOrder.CustomerName
	orderResponse.OrderedAt = entityOrder.OrderedAt.String()

	items, err := s.itemRepository.FindAllItemByOrderID(orderID)
	if err != nil {
		fmt.Println(err)
		return orderResponse, err
	}

	if len(items) > 0 {

		for _, item := range items {

			var orderItem dto.OrderItemResponse
			orderItem.ItemCode = item.ItemCode
			orderItem.Description = item.Description
			orderItem.Quantity = item.Quantity
			orderItem.ID = item.ID

			orderResponse.OrderItems = append(orderResponse.OrderItems, orderItem)

		}

	}

	return orderResponse, nil
}

func (s *orderService) Update(orderID int64, req *dto.OrderRequestUpdate) (dto.OrderResponse, error) {
	var orderResponse dto.OrderResponse

	entityOrder, err := s.orderRepository.FindOrderById(orderID)
	if err != nil {
		fmt.Println(err)
		err = errors.New("order tidak ditemukan")
		return orderResponse, err
	}

	_, err = s.orderRepository.UpdateOrder(orderID, req)
	if err != nil {
		fmt.Println(err)
		return orderResponse, err
	}

	// entityOrder, err = s.orderRepository.FindOrderById(orderID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return orderResponse, err
	// }

	// orderResponse.ID = entityOrder.ID
	// orderResponse.CustomerName = entityOrder.CustomerName
	// orderResponse.OrderedAt = entityOrder.OrderedAt.String()

	orderResponse, err = s.Read(entityOrder.ID)
	if err != nil {
		fmt.Println(err)
		return orderResponse, err
	}

	return orderResponse, nil
}

func (s *orderService) Delete(orderID int64) error {

	entityOrder, err := s.orderRepository.FindOrderById(orderID)
	if err != nil {
		fmt.Println(err)
		err = errors.New("order tidak ditemukan")
		return err
	}

	_, err = s.orderRepository.DeleteOrder(entityOrder.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *orderService) GetAll() ([]dto.OrderResponse, error) {

	var orders []dto.OrderResponse

	entityOrders, err := s.orderRepository.FindAllOrder()
	if err != nil {
		fmt.Println(err)
		return orders, err
	}

	for _, entityOrder := range entityOrders {

		var order dto.OrderResponse

		order.ID = entityOrder.ID
		order.CustomerName = entityOrder.CustomerName
		order.OrderedAt = entityOrder.OrderedAt.String()

		items, _ := s.itemRepository.FindAllItemByOrderID(order.ID)

		if len(items) > 0 {

			for _, item := range items {

				var orderItem dto.OrderItemResponse
				orderItem.ItemCode = item.ItemCode
				orderItem.Description = item.Description
				orderItem.Quantity = item.Quantity
				orderItem.ID = item.ID

				order.OrderItems = append(order.OrderItems, orderItem)

			}

		}

		orders = append(orders, order)

	}

	return orders, nil
}
