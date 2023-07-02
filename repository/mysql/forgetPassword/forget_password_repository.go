package forgetPassword

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewForgetPasswordRepository(db *gorm.DB) interfaces.ForgetPasswordRepository {
	return &repository{
		db: db,
	}
}

// FindOneByCode implements interfaces.ForgetPasswordRepository.
func (r *repository) FindOneByCode(code string) (*entity.ForgotPassword, error) {
	var forgotPassword entity.ForgotPassword
	err := r.db.Where("code = ?", code).First(&forgotPassword).Error
	if err != nil {
		return nil, err
	}

	return &forgotPassword, nil
}

// InsertForgetPassword implements interfaces.ForgetPasswordRepository.
func (r *repository) InsertForgetPassword(input *entity.ForgotPassword) (*entity.ForgotPassword, error) {
	err := r.db.Create(input).Error
	if err != nil {
		return nil, err
	}

	return input, nil
}

// UpdateForgetPassword implements interfaces.ForgetPasswordRepository.
func (r *repository) UpdateForgetPassword(input *entity.ForgotPassword) (*entity.ForgotPassword, error) {
	err := r.db.Save(input).Error
	if err != nil {
		return nil, err
	}

	return input, nil
}
