package road

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/consts"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
	"net/http"
	"time"
)

func (h *Handler) SaveRoadCard(c *gin.Context) {
	var request models.SaveRoadCardRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ContextError{
			ErrorMessage: consts.InvalidRequestData,
		})
		h.log.Error(consts.InvalidRequestData, logger.Error(err))
		return
	}

	response := make(chan models.SaveRoadCardResponse, 1)

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.serviceManager.IRoadService.SaveRoadCard(
		h.config.Manticore.IndexerUrl, h.config.Manticore.MergeUrl,
		request, response)

	select {
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, models.ContextError{
			ErrorMessage: consts.TimeLimitExceed,
		})
		h.log.Error(consts.TimeLimitExceed, logger.Error(err))
		return
	case result := <-response:
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, models.ContextError{
				ErrorMessage: consts.InvalidRequestData,
			})
			h.log.Error(result.Message, logger.Error(result.Error))
			return
		}

		c.JSON(http.StatusOK, result)
	}

}
