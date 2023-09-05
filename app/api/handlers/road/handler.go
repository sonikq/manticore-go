package road

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/services"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/configs"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/cache"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
)

type HandlerConfig struct {
	Conf           configs.Config
	Logger         logger.Logger
	Cache          *cache.Cache
	ServiceManager services.Service
}

type Handler struct {
	config         configs.Config
	log            logger.Logger
	cache          *cache.Cache
	serviceManager services.Service
}

func New(cfg *HandlerConfig) *Handler {
	return &Handler{
		config:         cfg.Conf,
		log:            cfg.Logger,
		cache:          cfg.Cache,
		serviceManager: cfg.ServiceManager,
	}
}
