package services

import (
	"golang.org/x/crypto/bcrypt"

	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/repository"
)

type AuthService struct {
	repo repository.IAuthRepo
}

func NewAuthService(repo repository.IAuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(request models.LoginRequest, response chan models.LoginResponse) {
	psswdHashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		response <- models.LoginResponse{
			UserId:  -1,
			Message: err.Error(),
			Error:   err,
		}
		return
	}

	user := models.LoginRequest{
		Username: request.Username,
		Password: string(psswdHashed),
	}

	result := s.repo.Login(user)

	response <- models.LoginResponse{
		UserId:  result.UserId,
		Message: result.Message,
		Error:   result.Error,
	}
}

func (s *AuthService) Logout(request models.LogoutRequest, response chan models.LogoutResponse) {
	result := s.repo.Logout(request)

	response <- models.LogoutResponse{
		Message: result.Message,
		Error:   result.Error,
	}

}
