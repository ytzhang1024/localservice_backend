package api

import (
	"6251/localservice/service"
	"6251/localservice/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderApi struct {
	BaseApi
	Service *service.OrderService
}

func NewOrderApi() OrderApi {
	return OrderApi{
		BaseApi: NewBaseApi(),
		Service: service.NewOrderService(),
	}
}

func (m OrderApi) AddOrder(c *gin.Context) {
	var iOrderAddDTO dto.OrderAddDTO
	iOrderAddDTO.Status = "Pending"
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iOrderAddDTO}).GetError(); err != nil {
		return
	}

	err := m.Service.AddOrder(&iOrderAddDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Order Adding Successful",
		Data: iOrderAddDTO,
	})
}

func (m OrderApi) GetOrderByUserId(c *gin.Context) {
	var iOrderDTO dto.OrderDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iOrderDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	OrderList, err := m.Service.GetOrderByUserId(&iOrderDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Getting User's Orders Successful",
		Data: OrderList,
	})
}

func (m OrderApi) GetOrderByServiceId(c *gin.Context) {
	var iOrderDTO dto.OrderDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iOrderDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	OrderList, err := m.Service.GetOrderByServiceId(&iOrderDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Getting service's Orders Successful",
		Data: OrderList,
	})
}

func (m OrderApi) UpdateOrder(c *gin.Context) {
	var iOrderUpdateDTO dto.OrderUpdateStateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iOrderUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	err := m.Service.UpdateOrder(&iOrderUpdateDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Update Order Successful",
	})
}

// 查找所有关于此用户的Update Info
