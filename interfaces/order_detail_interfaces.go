package interfaces

import "github.com/saufiroja/online-course-api/models/entity"

type OrderDetailRepository interface {
	InsertOrderDetail(input entity.OrderDetail) error
}

type OrderDetailService interface {
	InsertOrderDetail(input entity.OrderDetail) error
}
