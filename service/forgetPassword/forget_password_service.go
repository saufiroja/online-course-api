package forgetpassword

import (
	"time"

	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r       interfaces.ForgetPasswordRepository
	userSvc interfaces.UserService
	conf    *config.AppConfig
}

func NewForgetPasswordService(
	r interfaces.ForgetPasswordRepository,
	userSvc interfaces.UserService,
	conf *config.AppConfig,
) interfaces.ForgetPasswordService {
	return &service{
		r:       r,
		userSvc: userSvc,
		conf:    conf,
	}
}

// InsertForgetPassword implements interfaces.ForgetPasswordService.
func (s *service) InsertForgetPassword(input *dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, error) {
	user, err := s.userSvc.FindUserByEmail(input.Email)
	if err != nil {
		return nil, utils.HandlerError(404, "user not found")
	}

	if user == nil {
		return nil, utils.HandlerError(200, "success, please check your email to reset password")
	}

	dateTimes := time.Now().Add(time.Hour * 24)
	forgotPassword := &entity.ForgotPassword{
		UserID:    &user.ID,
		Valid:     true,
		Code:      utils.RandString(6),
		ExpiredAt: &dateTimes,
	}

	res, err := s.r.InsertForgetPassword(forgotPassword)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to insert forgot password")
	}

	data := dto.ForgotPasswordEmailRequestBody{
		SUBJECT: "Code Forgot Password",
		EMAIL:   user.Email,
		CODE:    forgotPassword.Code,
	}

	go utils.SendForgotPassword(user.Email, data, s.conf)

	return res, nil
}

// UpdateForgetPassword implements interfaces.ForgetPasswordService.
func (s *service) UpdateForgetPassword(input *dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, error) {
	code, err := s.r.FindOneByCode(input.Code)
	if err != nil {
		return nil, utils.HandlerError(404, "code not found")
	}

	user, err := s.userSvc.FindUserById(int(*code.UserID))
	if err != nil {
		return nil, utils.HandlerError(404, "user not found")
	}

	dataUser := &dto.UserUpdateRequestBody{
		Password: &input.Password,
	}

	_, err = s.userSvc.UpdateUser(int(user.ID), dataUser)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to update user")
	}

	code.Valid = false

	_, err = s.r.UpdateForgetPassword(code)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to update forgot password")
	}

	return code, nil
}
