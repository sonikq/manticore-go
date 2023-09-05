package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
)

type RoadDB struct {
	*sqlx.DB
}

type IRoadDB interface {
	AddToKillList(roadID int) error
	GenerateNewID() (int, error)
	AddToDelta(models.SaveRoadCardRequest) error
	CloseDB() error
}

func NewRoadDB(db *sqlx.DB) *RoadDB {
	return &RoadDB{
		DB: db,
	}
}

func (db *RoadDB) AddToKillList(roadID int) error {
	_, err := db.Exec(addToKillList, roadID)

	return err
}

func (db *RoadDB) GenerateNewID() (int, error) {
	var id int

	row := db.QueryRow(getNewID)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (db *RoadDB) AddToDelta(req models.SaveRoadCardRequest) error {
	_, err := db.Exec(addToDelta,
		req.ID, req.RoadID, req.ValueOfTheRoadGID, pq.Array(req.RegionGID),
		req.FullName, req.RoadNumberFull, req.EgradNumber, req.IsChecked,
		pq.Array(req.RoadCategoryGID), pq.Array(req.Capacity), pq.Array(req.SpeedLimit), req.JSONData,
	)

	return err
}

func (db *RoadDB) CloseDB() error {
	err := db.Close()

	return err
}
