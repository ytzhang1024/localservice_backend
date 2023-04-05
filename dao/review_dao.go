package dao

import (
	"6251/localservice/model"
	"6251/localservice/service/dto"
)

type ReviewDao struct {
	BaseDao
}

var reviewDao *ReviewDao

func NewReviewDao() *ReviewDao {
	if reviewDao == nil {
		reviewDao = &ReviewDao{NewBaseDao()}
	}

	return reviewDao
}

func (m *ReviewDao) AddReview(iReviewDTO *dto.ReviewDTO) error {
	var iReview model.Review
	iReviewDTO.ConvertToModel(&iReview)

	err := m.Orm.Save(&iReview).Error
	if err == nil {
		iReviewDTO.ID = iReview.ID
	}

	return err
}

func (m *ReviewDao) DeleteReviewById(id uint) error {
	return m.Orm.Delete(&model.Review{}, id).Error
}

func (m *ReviewDao) GetReviewByServiceId(id uint) ([]model.Review, error) {
	var ReviewList []model.Review
	err := m.Orm.Model(&model.Review{}).Where("service_id = ? ", id).Find(&ReviewList).Error
	return ReviewList, err
}
