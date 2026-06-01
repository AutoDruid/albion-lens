package fame

import (
	albionlens "github.com/AutoDruid/albion-lens"
	"github.com/AutoDruid/albion-lens/types/events"
	"github.com/AutoDruid/photon-parser"
)

type UpdateFameEvent struct {
	FameGained float64
	TotalFame  int64
}

func decodeUpdateFame(params []photon.ParameterV18) (UpdateFameEvent, bool) {
	var e UpdateFameEvent

	v, ok := params[1].IntValue()
	if !ok {
		return e, false
	}
	e.TotalFame = v

	v, ok = params[2].IntValue()
	if !ok {
		return e, false
	}
	e.FameGained = float64(v) / 10000

	return e, true
}

var UpdateFameHandler = albionlens.EventHandler[UpdateFameEvent]{
	Code:    events.UpdateFame,
	Handler: decodeUpdateFame,
}
