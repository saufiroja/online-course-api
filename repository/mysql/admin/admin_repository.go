package admin

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// CountAdmin implements interfaces.AdminRepository.
func (r *repository) CountAdmin() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Admin{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

// DeleteAdmin implements interfaces.AdminRepository.
func (r *repository) DeleteAdmin(admin entity.Admin) error {
	err := r.db.Delete(&admin).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAllAdmin implements interfaces.AdminRepository.
func (r *repository) FindAllAdmin(offset int, limit int) ([]entity.Admin, error) {
	var admins []entity.Admin
	err := r.db.Scopes(utils.Paginate(offset, limit)).Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return admins, nil
}

// FindOneAdminByEmail implements interfaces.AdminRepository.
func (r *repository) FindOneAdminByEmail(email string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// FindOneAdminByID implements interfaces.AdminRepository.
func (r *repository) FindOneAdminByID(id int) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// InsertAdmin implements interfaces.AdminRepository.
func (r *repository) InsertAdmin(admin entity.Admin) (*entity.Admin, error) {
	err := r.db.Create(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// UpdateAdmin implements interfaces.AdminRepository.
func (r *repository) UpdateAdmin(admin entity.Admin) (*entity.Admin, error) {
	err := r.db.Save(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &repository{db}
}
