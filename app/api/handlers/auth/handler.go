package auth

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/services"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/configs"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
)

type HandlerConfig struct {
	Conf           configs.Config
	Logger         logger.Logger
	ServiceManager services.Service
}

type Handler struct {
	config         configs.Config
	log            logger.Logger
	serviceManager services.Service
}

func New(cfg *HandlerConfig) *Handler {
	return &Handler{
		config:         cfg.Conf,
		log:            cfg.Logger,
		serviceManager: cfg.ServiceManager,
	}
}
