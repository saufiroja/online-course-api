package register

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	s    interfaces.UserService
	conf *config.AppConfig
}

func NewRegisterService(s interfaces.UserService, conf *config.AppConfig) interfaces.RegisterService {
	return &service{
		s,
		conf,
	}
}

// Register implements interfaces.RegisterService.
func (s *service) Register(input *dto.UserRequestBody) error {
	user, err := s.s.InsertUser(input)
	if err != nil {
		return utils.HandlerError(500, "failed to register user")
	}

	// send email
	data := dto.EmailVerification{
		SUBJECT:           "Email Verification",
		EMAIL:             user.Email,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go utils.SendVerification(input.Email, data, s.conf)

	return nil
}
