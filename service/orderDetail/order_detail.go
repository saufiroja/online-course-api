package orderdetail

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r interfaces.OrderDetailRepository
}

func NewOrderDetailService(r interfaces.OrderDetailRepository) interfaces.OrderDetailService {
	return &service{
		r: r,
	}
}

// InsertOrderDetail implements interfaces.OrderDetailService.
func (s *service) InsertOrderDetail(input entity.OrderDetail) error {
	err := s.r.InsertOrderDetail(input)
	if err != nil {
		return utils.HandlerError(400, err.Error())
	}

	return nil
}
