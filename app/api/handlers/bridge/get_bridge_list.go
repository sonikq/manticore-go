package bridge

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/utils"
	"io"
	"net/http"
)

func (h *Handler) GetBridgeList(c *gin.Context) {

	var request models.GetListRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		h.log.Error("Invalid request data: ", logger.Error(err))
		return
	}

	err := utils.ValidateIndex(h.config.Manticore.SearchUrl, request.Index)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid index"})
		h.log.Error("Error in validating index: ", logger.Error(err))
		return
	}

	reqBody := utils.LoadParams(models.GetListRequestLayout, request.Index, request.FullName)

	resp, err := http.Post(h.config.Manticore.SearchUrl, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in sending post request: ", logger.Error(err))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oops, something went wrong!"})
		h.log.Error("Error in reading body: ", logger.Error(err))
		return
	}

	result := utils.GetList(body)

	c.Data(http.StatusOK, "application/json", result)

}
