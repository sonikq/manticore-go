package repository

import (
	"bytes"
	"fmt"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/consts"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	cache2 "gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/cache"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/db"
	"net/http"
	"time"
)

type RoadRepo struct {
	cache *cache2.Cache
	db    *db.DB
}

func NewRoadRepo(cache *cache2.Cache, db *db.DB) *RoadRepo {
	return &RoadRepo{
		cache: cache,
		db:    db,
	}
}

func (r *RoadRepo) SaveRoadCard(idxUrl string, mergeUrl string, request models.SaveRoadCardRequest) models.SaveRoadCardResponse {

	start := time.Now()

	if msg, err := r.saveRoadCardInDB(request); err != nil {
		return models.SaveRoadCardResponse{
			Message: msg,
			Error:   err,
		}
	}

	if msg, err := mergeIndexes(idxUrl, mergeUrl); err != nil {
		return models.SaveRoadCardResponse{
			Message: msg,
			Error:   err,
		}
	}

	fmt.Println(time.Since(start))

	return models.SaveRoadCardResponse{
		Message: consts.SaveRoadCardSuccess,
		Error:   nil,
	}
}

func (r *RoadRepo) saveRoadCardInDB(request models.SaveRoadCardRequest) (string, error) {

	err := r.db.IRoadDB.AddToKillList(request.RoadID)
	if err != nil {
		return consts.AddKillListError, err
	}

	id, err := r.db.IRoadDB.GenerateNewID()
	if err != nil {
		return consts.GetNewIDError, err
	}

	request.ID = id

	err = r.db.IRoadDB.AddToDelta(request)
	if err != nil {
		return consts.AddToDeltaError, err
	}

	return "", nil
}

func mergeIndexes(idxUrl string, mergeUrl string) (string, error) {
	var buf bytes.Buffer

	s := []byte(`{"index":"idx_delta"}`)
	buf.Write(s)

	_, err := http.Post(idxUrl, "application/json", &buf)
	if err != nil {
		return consts.IdxDeltaError, err
	}

	buf.Reset()

	s = []byte(`{"main_index":"idx_main","delta_index":"idx_delta"}`)
	buf.Write(s)

	_, err = http.Post(idxUrl, "application/json", &buf)
	if err != nil {
		return consts.IdxMergeError, err
	}

	return "", nil
	//cmd := exec.Command(consts.Indexer, consts.IdxDelta, consts.Rotate)
	//
	//if err := cmd.Run(); err != nil {
	//	return consts.IdxDeltaError, err
	//}
	//
	//cmd = exec.Command(consts.Indexer, consts.Merge, consts.IdxMain, consts.IdxDelta, consts.Rotate)
	//
	//err := cmd.Run()
	//if err != nil {
	//	return consts.IdxMergeError, err
	//}
	//return "success", nil
}
