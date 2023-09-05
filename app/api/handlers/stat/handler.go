package stat

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/configs"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/cache"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
)

type HandlerConfig struct {
	Conf   configs.Config
	Logger logger.Logger
	Cache  *cache.Cache
}

type Handler struct {
	config configs.Config
	log    logger.Logger
	cache  *cache.Cache
}

func New(cfg *HandlerConfig) *Handler {
	return &Handler{
		config: cfg.Conf,
		log:    cfg.Logger,
		cache:  cfg.Cache,
	}
}
