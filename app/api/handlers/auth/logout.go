package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/consts"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/models"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
	"net/http"
	"time"
)

// Logout logs out users
// route /logout [post]
func (h *Handler) Logout(c *gin.Context) {
	var logout models.LogoutRequest

	err := c.BindJSON(&logout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ContextError{
			ErrorMessage: consts.InvalidRequestData,
		})
		h.log.Error(consts.InvalidRequestData, logger.Error(err))
		return
	}

	response := make(chan models.LogoutResponse, 1)

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.serviceManager.IAuthService.Logout(logout, response)

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
				ErrorMessage: consts.InternalServerError,
			})
			h.log.Error(result.Message, logger.Error(result.Error))
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
