package handler

import (
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/accident"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/auth"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/bridge"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/common"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/road"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/api/handlers/stat"
)

type Handler struct {
	AuthHandler     auth.Handler
	RoadHandler     road.Handler
	CommonHandler   common.Handler
	StatHandler     stat.Handler
	AccidentHandler accident.Handler
	BridgeHandler   bridge.Handler
}
