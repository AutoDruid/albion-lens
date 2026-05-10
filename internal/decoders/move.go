package decoders

import (
	"github.com/AutoDruid/albion-lens/models"
	"github.com/AutoDruid/photon-parser"
)

const (
	moveParamEntityID  uint8 = 0
	moveParamPos       uint8 = 1
	moveParamRotation  uint8 = 2
	moveParamTargetPos uint8 = 3
	moveParamSpeed     uint8 = 4
)

func DecodeMove(params []photon.ParameterV18) (models.MoveEvent, bool) {
	var e models.MoveEvent

	for _, p := range params {
		switch p.ID() {
		case moveParamEntityID:
			v, ok := p.IntValue()
			if !ok {
				return e, false
			}
			e.EntityID = uint64(v)

		case moveParamPos:
			var i int
			for _, v := range p.Float32ArrayValue() {
				switch i {
				case 0:
					e.Pos.X = v
				case 1:
					e.Pos.Z = v
				}
				i++
			}

		case moveParamRotation:
			v, ok := p.Float32Value()
			if !ok {
				return e, false
			}
			e.Rotation = v

		case moveParamTargetPos:
			var i int
			for _, v := range p.Float32ArrayValue() {
				switch i {
				case 0:
					e.TargetPos.X = v
				case 1:
					e.TargetPos.Z = v
				}
				i++
			}

		case moveParamSpeed:
			v, ok := p.Float32Value()
			if !ok {
				return e, false
			}
			e.Speed = v
		}
	}

	return e, e.EntityID != 0
}
