package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/accident"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/auth"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/bridge"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/common"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/road"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/stat"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/services"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/configs"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/cache"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/logger"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/pkg/middlewares"
)

type Option struct {
	Conf           configs.Config
	Logger         logger.Logger
	ServiceManager services.Service
	Cache          *cache.Cache
}

func New(option Option) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	h := &handler.Handler{
		AuthHandler: *auth.New(&auth.HandlerConfig{
			Conf:           option.Conf,
			Logger:         option.Logger,
			ServiceManager: option.ServiceManager,
		}),
		RoadHandler: *road.New(&road.HandlerConfig{
			Conf:           option.Conf,
			Logger:         option.Logger,
			Cache:          option.Cache,
			ServiceManager: option.ServiceManager,
		}),
		CommonHandler: *common.New(&common.HandlerConfig{
			Conf:   option.Conf,
			Logger: option.Logger,
			Cache:  option.Cache,
		}),
		StatHandler: *stat.New(&stat.HandlerConfig{
			Conf:   option.Conf,
			Logger: option.Logger,
			Cache:  option.Cache,
		}),
		AccidentHandler: *accident.New(&accident.HandlerConfig{
			Conf:   option.Conf,
			Logger: option.Logger,
			Cache:  option.Cache,
		}),
		BridgeHandler: *bridge.New(&bridge.HandlerConfig{
			Conf:   option.Conf,
			Logger: option.Logger,
			Cache:  option.Cache,
		}),
	}

	// no auth required endpoints
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Geogracom!",
		})
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	//Authentication endpoints
	_auth := router.Group("/auth")
	{
		_auth.POST("/login", h.AuthHandler.Login)
		_auth.POST("/logout", h.AuthHandler.Logout, middlewares.IsAuthorized())
	}

	// Road endpoints
	_road := router.Group("/road")
	_road.Use(middlewares.IsAuthorized())
	{
		_road.POST("/getlist", h.RoadHandler.GetRoadList)
		_road.POST("/getcard", h.RoadHandler.GetRoadCard)
		_road.POST("/getfeatures", h.RoadHandler.GetRoadFeatures)
		_road.POST("/savecard", h.RoadHandler.SaveRoadCard)
	}

	// Common endpoints
	_common := router.Group("/common")
	_common.Use(middlewares.IsAuthorized())
	{
		_common.POST("/getlist", h.CommonHandler.GetList)
		_common.POST("/getcard", h.CommonHandler.GetCard)
		_common.POST("/getfeatures", h.CommonHandler.GetFeatures)
	}

	// Accident endpoints
	_accident := router.Group("/accident")
	_accident.Use(middlewares.IsAuthorized())
	{
		_accident.POST("/getlist", h.AccidentHandler.GetAccidentList)
		_accident.POST("/getcard", h.AccidentHandler.GetAccidentCard)
		_accident.POST("/getfeatures", h.AccidentHandler.GetAccidentFeatures)
	}

	// Bridge endpoints
	_bridge := router.Group("/bridge")
	_bridge.Use(middlewares.IsAuthorized())
	{
		_bridge.POST("/getlist", h.BridgeHandler.GetBridgeList)
		_bridge.POST("/getcard", h.BridgeHandler.GetBridgeCard)
		_bridge.POST("/getfeatures", h.BridgeHandler.GetBridgeFeatures)
	}

	_stat := router.Group("/stat")
	_stat.Use(middlewares.IsAuthorized())
	{
		_stat.POST("/getform", h.StatHandler.GetForm)
		_stat.POST("/getformlist", h.StatHandler.GetFormList)
		_stat.POST("/downloadform", h.StatHandler.DownloadForm)
		_stat.POST("/parseform", h.StatHandler.ParseForm)
		_stat.POST("/uploadform", h.StatHandler.UploadForm)
		_stat.POST("/setformvalue", h.StatHandler.SetFormValue)

	}

	return router
}
