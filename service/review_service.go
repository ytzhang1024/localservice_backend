package service

import (
	"6251/localservice/dao"
	"6251/localservice/model"
	"6251/localservice/service/dto"
)

type ReviewService struct {
	BaseService
	Dao *dao.ReviewDao
}

var reviewService *ReviewService

func NewReviewService() *ReviewService {
	if reviewService == nil {
		reviewService = &ReviewService{
			Dao: dao.NewReviewDao(),
		}
	}
	return reviewService
}

func (m *ReviewService) AddReview(iReviewAddDTO *dto.ReviewDTO) error {
	return m.Dao.AddReview(iReviewAddDTO)
}

func (m *ReviewService) DeleteReviewById(iCommonIDDTO *dto.CommonIDDTO) error {
	return m.Dao.DeleteReviewById(iCommonIDDTO.ID)
}

func (m *ReviewService) GetReviewByServiceId(iReviewDTO *dto.ReviewDTO) ([]model.Review, error) {
	return m.Dao.GetReviewByServiceId(iReviewDTO.ServiceId)
}
