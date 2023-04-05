package dao

import (
	"6251/localservice/model"
	"6251/localservice/service/dto"
)

type OrderDao struct {
	BaseDao
}

var orderDao *OrderDao

func NewOrderDao() *OrderDao {
	if orderDao == nil {
		orderDao = &OrderDao{NewBaseDao()}
	}

	return orderDao
}

func (m *OrderDao) AddOrder(iOrderAddDTO *dto.OrderAddDTO) error {
	var iOrderAdd model.Order
	iOrderAddDTO.ConvertToModel(&iOrderAdd)

	err := m.Orm.Save(&iOrderAdd).Error
	if err == nil {
		iOrderAddDTO.ID = iOrderAdd.ID
	}

	return err
}

func (m *OrderDao) DeleteOrderById(id uint) error {
	return m.Orm.Delete(&model.Order{}, id).Error
}

// user can see their orders
func (m *OrderDao) GetOrderByUserId(id uint) ([]model.Order, error) {
	var OrderList []model.Order
	err := m.Orm.Model(&model.Order{}).Where("customer_id = ? ", id).Find(&OrderList).Error
	return OrderList, err
}

// service provider can view their orders
func (m *OrderDao) GetOrderByServiceId(id uint) ([]model.Order, error) {
	var OrderList []model.Order
	err := m.Orm.Model(&model.Order{}).Where("service_id = ? ", id).Find(&OrderList).Error
	return OrderList, err
}

func (m *OrderDao) UpdateOrder(iOrderUpdateStateDTO *dto.OrderUpdateStateDTO) error {
	var iOrder model.Order

	m.Orm.First(&iOrder, iOrderUpdateStateDTO.ID)
	iOrderUpdateStateDTO.ConvertToModel(&iOrder)

	return m.Orm.Save(&iOrder).Error
}
