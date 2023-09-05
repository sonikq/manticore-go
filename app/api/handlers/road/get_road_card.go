package road

import (
	"bytes"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/consts"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/utils"
)

func (h *Handler) GetRoadCard(c *gin.Context) {

	var request models.GetCardRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ContextError{
			ErrorMessage: consts.InvalidRequestData,
		})
		h.log.Error(consts.InvalidRequestData, logger.Error(err))
		return
	}

	err := utils.ValidateIndex(h.config.Manticore.SearchUrl, request.Index)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ContextError{
			ErrorMessage: consts.InternalServerError,
		})
		h.log.Error(consts.IndexNotFound, logger.Error(err))
		return
	}

	_, err = h.cache.Get(request.ID)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ContextError{
			ErrorMessage: consts.NotAuthorized,
		})
		h.log.Error(consts.NotAuthorized, logger.Error(err))
		return
	}

	reqBody := utils.LoadParams(models.GetCardRequestLayout, request.Index, request.ID)

	resp, err := http.Post(h.config.Manticore.SearchUrl, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ContextError{
			ErrorMessage: consts.InternalServerError,
		})
		h.log.Error(consts.HTTPRequestError, logger.Error(err))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ContextError{
			ErrorMessage: consts.InternalServerError,
		})
		h.log.Error(consts.ReadBodyError, logger.Error(err))
		return
	}

	result := utils.GetCard(body)

	c.Data(http.StatusOK, "application/json", result)
}
