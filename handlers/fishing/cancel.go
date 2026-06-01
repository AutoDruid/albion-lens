package fishing

import (
	albionlens "github.com/AutoDruid/albion-lens"
	"github.com/AutoDruid/albion-lens/types/events"
	"github.com/AutoDruid/photon-parser"
)

type FishingCancelEvent struct {
	PlayerID     uint64
	CancelReason uint8
}

func decodeFishingCancel(params []photon.ParameterV18) (FishingCancelEvent, bool) {
	var e FishingCancelEvent

	v, ok := params[0].IntValue()
	if !ok {
		return e, false
	}
	e.PlayerID = uint64(v)

	v, ok = params[1].IntValue()
	if !ok {
		return e, false
	}
	e.CancelReason = uint8(v)

	return e, true
}

var FishingCancelHandler = albionlens.EventHandler[FishingCancelEvent]{
	Code:    events.FishingCancel,
	Handler: decodeFishingCancel,
}
