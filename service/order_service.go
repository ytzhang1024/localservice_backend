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

func (m *OrderService) AddOrder(iOrderAddDTO *dto.OrderDTO) error {
	return m.Dao.AddOrder(iOrderAddDTO)
}

func (m *OrderService) GetOrderByUserId(iOrderDTO *dto.OrderDTO) ([]model.Order, error) {
	return m.Dao.GetOrderByUserId(iOrderDTO.CustomerId)
}

func (m *OrderService) GetOrderById(id uint) ([]model.Order, error) {
	return m.Dao.GetOrderById(uint(id))
}

func (m *OrderService) GetOrderByServiceId(iOrderDTO *dto.OrderDTO) ([]model.Order, error) {
	return m.Dao.GetOrderByServiceId(iOrderDTO.ServiceId)
}

func (m *OrderService) GetOrderByProviderId(iOrderDTO *dto.OrderDTO) ([]model.Order, error) {
	return m.Dao.GetOrderByProviderId(iOrderDTO.ProviderId)
}

func (m *OrderService) UpdateOrder(iOrderDTO *dto.OrderUpdateStatusDTO) error {
	if iOrderDTO.ID == 0 {
		return errors.New("Invalid Order ID")
	}

	return m.Dao.UpdateOrder(iOrderDTO)
}

func (m *OrderService) UpdateOrderStatus(iCommonIDDTO *dto.CommonIDDTO, status string) error {
	return m.Dao.UpdateOrderStatus(iCommonIDDTO.ID, status)
}
