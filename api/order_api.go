package api

import (
	"6251/localservice/service"
	"6251/localservice/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	var iOrderAddDTO dto.OrderDTO
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

func (m OrderApi) GetOrderByProviderId(c *gin.Context) {
	var iOrderDTO dto.OrderDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iOrderDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	OrderList, err := m.Service.GetOrderByProviderId(&iOrderDTO)

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

	var iOrderDTO dto.OrderUpdateStatusDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iOrderDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	err := m.Service.UpdateOrder(&iOrderDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Update User Successful",
		Data: iOrderDTO,
	})
}

func (m OrderApi) GetOrderById(c *gin.Context) {
	var iOrderDTO dto.OrderDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iOrderDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	id_str := c.Query("order_id")
	id, _ := strconv.Atoi(id_str)

	OrderList, err := m.Service.GetOrderById(uint(id))

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Getting Order Successful",
		Data: OrderList,
	})
}

func (m OrderApi) UpdateOrderStatus(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	status := c.Query("status")
	println(status)

	err := m.Service.UpdateOrderStatus(&iCommonIDDTO, status)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Update Order Status Successful",
	})
}
