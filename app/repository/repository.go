package repository

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	cache2 "gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/cache"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/db"
)

type IAuthRepo interface {
	Login(request models.LoginRequest) models.LoginResponse
	Logout(request models.LogoutRequest) models.LogoutResponse
	LoadAuthMatrix(userId int) error
	UnloadAuthMatrix()
}

type IRoadRepo interface {
	SaveRoadCard(idxUrl string, mergeUrl string, request models.SaveRoadCardRequest) models.SaveRoadCardResponse
	saveRoadCardInDB(request models.SaveRoadCardRequest) (string, error)
}

type Repository struct {
	IAuthRepo
	IRoadRepo
}

func NewRepository(cache *cache2.Cache, db *db.DB) *Repository {
	return &Repository{
		IAuthRepo: NewAuthRepo(cache, db),
		IRoadRepo: NewRoadRepo(cache, db),
	}
}
