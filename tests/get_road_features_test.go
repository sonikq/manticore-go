package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/utils"
)

const (
	url = "http://10.10.10.72:9308/json/search"
)

func TestGetRoadFeatures(t *testing.T) {

	request := models.GetFeaturesRequest{
		EditMode:    0,
		Zoom:        11,
		ScaleFactor: 1,
		Points: []float64{
			4115904.1939143892,
			7438170.934858386,
			4242483.912754641,
			7494122.839563134,
		},
	}

	reqBody := utils.LoadWFSParams(
		models.GetFeaturesRequestLayout,
		"idx_wfs_road_1000", 10000, 10000,
		request.Zoom,
		request.ScaleFactor,
		request.Points,
	)

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	assert.Equal(t, err, nil)

	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, err, nil)

	result := utils.GetFeatures(body)

	assert.Equal(t, true, json.Valid(result))
}

func TestGetList(t *testing.T) {

	request := models.GetListRequest{
		Index:    "idx_road_list",
		FullName: "Подъезд",
	}

	reqBody := utils.LoadParams(models.GetListRequestLayout, request.Index, request.FullName)

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	assert.Equal(t, err, nil)

	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, err, nil)

	result := utils.GetList(body)

	assert.Equal(t, true, json.Valid(result))
}

func TestGetCard(t *testing.T) {

	request := models.GetCardRequest{
		Index: "idx_road_card",
		ID:    87265,
	}

	reqBody := utils.LoadParams(models.GetCardRequestLayout, request.Index, request.ID)

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	assert.Equal(t, err, nil)

	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, err, nil)

	result := utils.GetCard(body)

	assert.Equal(t, true, json.Valid(result))
}
