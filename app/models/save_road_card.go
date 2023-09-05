package models

import (
	"github.com/jmoiron/sqlx/types"
)

type SaveRoadCardRequest struct {
	ID                int            `db:"id"`
	RoadID            int            `json:"road_id" db:"road_id"`
	ValueOfTheRoadGID int            `json:"value_of_the_road_gid" db:"value_of_the_road_gid"`
	RegionGID         []int          `json:"region_gid" db:"region_gid"`
	FullName          string         `json:"full_name" db:"full_name"`
	RoadNumberFull    string         `json:"road_number_full" db:"road_number_full"`
	EgradNumber       string         `json:"egrad_number" db:"egrad_number"`
	IsChecked         bool           `json:"is_checked" db:"is_checked"`
	RoadCategoryGID   []int          `json:"road_category_gid" db:"road_category_gid"`
	Capacity          []int          `json:"capacity" db:"capacity"`
	SpeedLimit        []int          `json:"speed_limit" db:"speed_limit"`
	JSONData          types.JSONText `json:"json_data" db:"json_data"`
}

type SaveRoadCardResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
