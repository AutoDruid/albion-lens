package fishing

import (
	albionlens "github.com/AutoDruid/albion-lens"
	"github.com/AutoDruid/albion-lens/types/events"
	"github.com/AutoDruid/photon-parser"
)

type FishingSubType uint8

const (
	FishingCast        FishingSubType = 3
	FishingCastEnd     FishingSubType = 4
	FishingBite        FishingSubType = 5
	FishingPulling     FishingSubType = 7
	FishingStopPulling FishingSubType = 8
	FishingCatch       FishingSubType = 9
	FishingFailed      FishingSubType = 10
	FishingNotAllowed  FishingSubType = 11
	FishingCancel      FishingSubType = 15
	// 14 observed — validate against traffic
)

type FishingStartEvent struct {
	PlayerID     uint64
	BobberID     uint64
	FishingRodID uint16
	SubType      FishingSubType
	ServerTicks  int64
}

func decodeFishingStart(params []photon.ParameterV18) (FishingStartEvent, bool) {
	var e FishingStartEvent

	v, ok := params[0].IntValue()
	if !ok {
		return e, false
	}
	e.PlayerID = uint64(v)

	v, ok = params[1].IntValue()
	if !ok {
		return e, false
	}
	e.BobberID = uint64(v)

	v, ok = params[2].IntValue()
	if !ok {
		return e, false
	}
	e.FishingRodID = uint16(v)

	v, ok = params[3].IntValue()
	if !ok {
		return e, false
	}
	e.SubType = FishingSubType(v)

	v, ok = params[4].IntValue()
	if !ok {
		return e, false
	}
	e.ServerTicks = v

	return e, true
}

var FishingStartHandler = albionlens.EventHandler[FishingStartEvent]{
	Code:    events.FishingStart,
	Handler: decodeFishingStart,
}
