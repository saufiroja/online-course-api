package user

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &repository{db}
}

// InsertUser implements interfaces.UserRepository.
func (r *repository) InsertUser(user *entity.User) (*entity.User, error) {
	err := r.db.Model(&entity.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByEmail implements interfaces.UserRepository.
func (r *repository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Model(&entity.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CountUser implements interfaces.UserRepository.
func (r *repository) CountUser() (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

// DeleteUser implements interfaces.UserRepository.
func (r *repository) DeleteUser(user *entity.User) error {
	err := r.db.Model(&entity.User{}).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAllUser implements interfaces.UserRepository.
func (r *repository) FindAllUser(offset int, limit int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Model(&entity.User{}).Scopes(utils.Paginate(offset, limit)).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

// FindUserByCodeVerified implements interfaces.UserRepository.
func (r *repository) FindUserByCodeVerified(codeVerified string) (*entity.User, error) {
	var user entity.User
	err := r.db.Model(&entity.User{}).Where("code_verified = ?", codeVerified).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserById implements interfaces.UserRepository.
func (r *repository) FindUserById(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Model(&entity.User{}).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser implements interfaces.UserRepository.
func (r *repository) UpdateUser(user *entity.User) (*entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
