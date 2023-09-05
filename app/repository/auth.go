package repository

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	cache2 "gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/cache"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/db"
	"math/rand"
	"time"
)

type AuthRepo struct {
	cache *cache2.Cache
	db    *db.DB
}

func NewAuthRepo(cache *cache2.Cache, db *db.DB) *AuthRepo {
	return &AuthRepo{
		cache: cache,
		db:    db,
	}
}

func (r *AuthRepo) Login(request models.LoginRequest) models.LoginResponse {
	id := rand.Intn(999999)

	if err := r.LoadAuthMatrix(id); err != nil {
		return models.LoginResponse{
			UserId:  0,
			Message: "Cannot load permissions matrix",
			Error:   err,
		}
	}

	return models.LoginResponse{
		UserId:  id,
		Message: "Successfully logged in",
		Error:   nil,
	}

}

func (r *AuthRepo) Logout(request models.LogoutRequest) models.LogoutResponse {

	r.UnloadAuthMatrix()

	return models.LogoutResponse{
		Message: "Successfully logged out",
		Error:   nil,
	}

}

func (r *AuthRepo) LoadAuthMatrix(userId int) error {

	rows, err := r.db.IAuthDB.LoadAuthMatrix(userId)
	if err != nil {
		return err
	}

	for rows.Next() {
		var id, lvl, code int

		err = rows.Scan(&id, &lvl, &code)
		if err != nil {
			return err
		}

		perms := models.Perms{
			Access: lvl,
			Code:   code,
		}

		r.cache.Set(id, perms, 180*time.Second)

	}

	return nil

}

func (r *AuthRepo) UnloadAuthMatrix() {

	r.cache.FlushCache()
}
