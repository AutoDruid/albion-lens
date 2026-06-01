package move

import (
	albionlens "github.com/AutoDruid/albion-lens"
	"github.com/AutoDruid/albion-lens/handlers"
	"github.com/AutoDruid/albion-lens/types/events"
	"github.com/AutoDruid/photon-parser"
)

type MoveEvent struct {
	EntityID  uint64
	Pos       handlers.Vec2
	TargetPos handlers.Vec2
	Rotation  float32
	Speed     float32
}

func decodeMove(params []photon.ParameterV18) (MoveEvent, bool) {
	var e MoveEvent

	v, ok := params[0].IntValue()
	if !ok {
		return e, false
	}
	e.EntityID = uint64(v)

	var i int
	for _, v := range params[1].Float32ArrayValue() {
		switch i {
		case 0:
			e.Pos.X = v
		case 1:
			e.Pos.Y = v
		}
		i++
	}

	rotation, ok := params[2].Float32Value()
	if !ok {
		return e, false
	}
	e.Rotation = rotation

	i = 0
	for _, v := range params[3].Float32ArrayValue() {
		switch i {
		case 0:
			e.TargetPos.X = v
		case 1:
			e.TargetPos.Y = v
		}
		i++
	}

	speed, ok := params[4].Float32Value()
	if !ok {
		return e, false
	}
	e.Speed = speed

	return e, e.EntityID != 0
}

var MoveHandler = albionlens.EventHandler[MoveEvent]{
	Code:    events.Move,
	Handler: decodeMove,
}
