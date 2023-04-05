package api

import (
	"6251/localservice/service"
	"6251/localservice/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReviewApi struct {
	BaseApi
	Service *service.ReviewService
}

func NewReviewApi() ReviewApi {
	return ReviewApi{
		BaseApi: NewBaseApi(),
		Service: service.NewReviewService(),
	}
}

func (m ReviewApi) AddReview(c *gin.Context) {
	var iReviewDTO dto.ReviewDTO

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iReviewDTO}).GetError(); err != nil {
		return
	}

	err := m.Service.AddReview(&iReviewDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Review Adding Successful",
		Data: iReviewDTO,
	})
}

func (m ReviewApi) DeleteReviewById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteReviewById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Delete Review Successful",
	})
}

func (m ReviewApi) GetReviewByServiceId(c *gin.Context) {
	var iReviewDTO dto.ReviewDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iReviewDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	ReviewList, err := m.Service.GetReviewByServiceId(&iReviewDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: http.StatusOK,
		Msg:  "Getting Reviews Successful",
		Data: ReviewList,
	})
}
