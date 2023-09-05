package services

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/repository"
)

type IAuthService interface {
	Login(request models.LoginRequest, response chan models.LoginResponse)
	Logout(request models.LogoutRequest, response chan models.LogoutResponse)
}

type IRoadService interface {
	SaveRoadCard(idxUrl string, mergeUrl string, request models.SaveRoadCardRequest, response chan models.SaveRoadCardResponse)
}

type Service struct {
	IAuthService
	IRoadService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IAuthService: NewAuthService(repos.IAuthRepo),
		IRoadService: NewRoadService(repos.IRoadRepo),
	}
}
