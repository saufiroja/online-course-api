package user

import (
	"fmt"

	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) interfaces.UserService {
	return &service{r}
}

// InsertUser implements interfaces.UserService.
func (s *service) InsertUser(input *dto.UserRequestBody) (*entity.User, error) {
	user, err := s.r.FindUserByEmail(input.Email)
	if err != nil && err.Error() != "record not found" {
		return nil, utils.HandlerError(404, "email already exists")
	}

	if user != nil {
		return nil, utils.HandlerError(404, "email already exists")
	}

	// hash password
	hashedPassword := utils.HashPassword(input.Password)

	users := &entity.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     hashedPassword,
		CodeVerified: utils.RandString(6),
	}

	if input.CreatedBy != nil {
		users.CreatedByID = input.CreatedBy
	}

	fmt.Println("users", users)
	res, err := s.r.InsertUser(users)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to register user")
	}

	return res, nil
}

// CountUser implements interfaces.UserService.
func (s *service) CountUser() (int64, error) {
	count, err := s.r.CountUser()
	if err != nil {
		return 0, utils.HandlerError(500, "failed to count user")
	}

	return count, nil
}

// DeleteUser implements interfaces.UserService.
func (s *service) DeleteUser(id int) error {
	// check user exist
	user, err := s.r.FindUserById(id)
	if err != nil {
		return utils.HandlerError(404, "user not found")
	}

	// delete user
	return s.r.DeleteUser(user)
}

// FindAllUser implements interfaces.UserService.
func (s *service) FindAllUser(offset int, limit int) ([]entity.User, error) {
	users, err := s.r.FindAllUser(offset, limit)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to get all user")
	}

	return users, nil
}

// FindUserByCodeVerified implements interfaces.UserService.
func (s *service) FindUserByCodeVerified(codeVerified string) (*entity.User, error) {
	users, err := s.r.FindUserByCodeVerified(codeVerified)
	if err != nil {
		return nil, utils.HandlerError(404, "user not found")
	}

	return users, nil
}

// FindUserByEmail implements interfaces.UserService.
func (s *service) FindUserByEmail(email string) (*entity.User, error) {
	users, err := s.r.FindUserByEmail(email)
	if err != nil {
		return nil, utils.HandlerError(404, "user not found")
	}

	return users, nil
}

// FindUserById implements interfaces.UserService.
func (s *service) FindUserById(id int) (*entity.User, error) {
	users, err := s.r.FindUserById(id)
	if err != nil {
		return nil, utils.HandlerError(404, "user not found")
	}

	return users, nil
}

// UpdateUser implements interfaces.UserService.
func (s *service) UpdateUser(id int, input *dto.UserUpdateRequestBody) (*entity.User, error) {
	user, err := s.r.FindUserById(id)
	if err != nil {
		return nil, utils.HandlerError(404, "user not found")
	}

	if input.Email != nil {
		if user.Email != *input.Email {
			user.Email = *input.Email
		}
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Password != nil {
		hashedPassword := utils.HashPassword(*input.Password)
		user.Password = hashedPassword
	}

	if input.UpdatedBy != nil {
		user.UpdatedByID = input.UpdatedBy
	}

	if input.EmailVerifiedAt != nil {
		user.EmailVerifiedAt = input.EmailVerifiedAt
	}

	updateUser, err := s.r.UpdateUser(user)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to update user")
	}

	return updateUser, nil
}
