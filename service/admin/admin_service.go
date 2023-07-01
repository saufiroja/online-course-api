package admin

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r interfaces.AdminRepository
}

func NewAdminService(s interfaces.AdminRepository) interfaces.AdminService {
	return &service{s}
}

// CountAdmin implements interfaces.AdminService.
func (s *service) CountAdmin() (int64, error) {
	return s.r.CountAdmin()
}

// DeleteAdmin implements interfaces.AdminService.
func (s *service) DeleteAdmin(id int) error {
	admin, err := s.r.FindOneAdminByID(id)
	if err != nil {
		return utils.HandlerError(404, "admin not found")
	}

	err = s.r.DeleteAdmin(*admin)
	if err != nil {
		return utils.HandlerError(500, "failed to delete admin")
	}

	return nil
}

// FindAllAdmin implements interfaces.AdminService.
func (s *service) FindAllAdmin(offset int, limit int) ([]entity.Admin, error) {
	res, err := s.r.FindAllAdmin(offset, limit)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to get admin")
	}

	return res, nil
}

// FindOneAdminByEmail implements interfaces.AdminService.
func (s *service) FindOneAdminByEmail(email string) (*entity.Admin, error) {
	admin, err := s.r.FindOneAdminByEmail(email)
	if err != nil {
		return nil, utils.HandlerError(404, "admin not found")
	}

	return admin, nil
}

// FindOneAdminByID implements interfaces.AdminService.
func (s *service) FindOneAdminByID(id int) (*entity.Admin, error) {
	admin, err := s.r.FindOneAdminByID(id)
	if err != nil {
		return nil, utils.HandlerError(404, "admin not found")
	}

	return admin, nil
}

// InsertAdmin implements interfaces.AdminService.
func (s *service) InsertAdmin(admin dto.AdminRequestBody) (*entity.Admin, error) {
	hash := utils.HashPassword(*admin.Password)
	data := entity.Admin{
		Email:    admin.Email,
		Name:     admin.Name,
		Password: hash,
	}

	if admin.CreatedBy != nil {
		data.CreatedByID = admin.CreatedBy
	}

	res, err := s.r.InsertAdmin(data)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to insert admin")
	}

	return res, nil
}

// UpdateAdmin implements interfaces.AdminService.
func (s *service) UpdateAdmin(id int, admin dto.AdminRequestBody) (*entity.Admin, error) {
	admins, err := s.r.FindOneAdminByID(id)
	if err != nil {
		return nil, utils.HandlerError(404, "admin not found")
	}

	admins.Name = admin.Name
	admins.Email = admin.Email

	if admin.Password != nil {
		hash := utils.HashPassword(*admin.Password)
		admins.Password = hash
	}

	if admin.UpdatedBy != nil {
		admins.UpdatedByID = admin.UpdatedBy
	}

	res, err := s.r.UpdateAdmin(*admins)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to update admin")
	}

	return res, nil
}
