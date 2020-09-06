package service

import (
	"database/sql"
	"go-practice/electricity-project/datamodels"
	"go-practice/electricity-project/repositories"
)

type IOrderService interface {
	GetOrderById(orderId int64) (order *datamodels.Order, err error)
	DeleteOrderById(orderId int64) bool
	UpdateOrder(order *datamodels.Order) error
	InsertOrder(order *datamodels.Order) (orderId int64, err error)
	GetAllOrder() (orders []*datamodels.Order, err error)
	GetAllOrderInfo() (orderInfos map[int]map[string]string, err error)
}

func NewOrderService(db *sql.DB) IOrderService {
	orderRepo := repositories.NewOrderManagerRepository("order", db)

	return &OrderService{orderRepository: orderRepo}
}

type OrderService struct {
	orderRepository repositories.IOrderRepository
}

func (o OrderService) GetOrderById(orderId int64) (order *datamodels.Order, err error) {
	return o.orderRepository.SelectByKey(orderId)
}

func (o OrderService) DeleteOrderById(orderId int64) bool {
	return o.orderRepository.Delete(orderId)
}

func (o OrderService) UpdateOrder(order *datamodels.Order) error {
	return o.orderRepository.Update(order)
}

func (o OrderService) InsertOrder(order *datamodels.Order) (orderId int64, err error) {
	return o.orderRepository.Insert(order)
}

func (o OrderService) GetAllOrder() (orders []*datamodels.Order, err error) {
	return o.orderRepository.SelectAll()
}

func (o OrderService) GetAllOrderInfo() (orderInfos map[int]map[string]string, err error) {
	return o.orderRepository.SelectAllWithInfo()
}
