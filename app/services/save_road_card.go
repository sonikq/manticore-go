package services

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/repository"
)

type RoadService struct {
	repo repository.IRoadRepo
}

func NewRoadService(repo repository.IRoadRepo) *RoadService {
	return &RoadService{repo: repo}
}

func (s *RoadService) SaveRoadCard(idxUrl string, mergeUrl string, request models.SaveRoadCardRequest, response chan models.SaveRoadCardResponse) {

	result := s.repo.SaveRoadCard(idxUrl, mergeUrl, request)

	response <- models.SaveRoadCardResponse{
		Message: result.Message,
		Error:   result.Error,
	}
}
