package service

import (
	"6251/localservice/dao"
	"6251/localservice/model"
	"6251/localservice/service/dto"
	"errors"
)

var orderService *OrderService

type OrderService struct {
	BaseService
	Dao *dao.OrderDao
}

func NewOrderService() *OrderService {
	if orderService == nil {
		orderService = &OrderService{
			Dao: dao.NewOrderDao(),
		}
	}

	return orderService
}

func (m *OrderService) AddOrder(iOrderAddDTO *dto.OrderAddDTO) error {
	return m.Dao.AddOrder(iOrderAddDTO)
}

func (m *OrderService) GetOrderByUserId(iOrderDTO *dto.OrderDTO) ([]model.Order, error) {
	return m.Dao.GetOrderByUserId(iOrderDTO.CustomerId)
}

func (m *OrderService) GetOrderByServiceId(iOrderDTO *dto.OrderDTO) ([]model.Order, error) {
	return m.Dao.GetOrderByServiceId(iOrderDTO.ServiceId)
}

func (m *OrderService) UpdateOrder(iOrderUpdateStateDTO *dto.OrderUpdateStateDTO) error {
	if iOrderUpdateStateDTO.ID == 0 {
		return errors.New("Invalid Order ID")
	}

	return m.Dao.UpdateOrder(iOrderUpdateStateDTO)
}
