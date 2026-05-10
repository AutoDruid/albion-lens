package decoders

import (
	"log"

	"github.com/AutoDruid/albion-lens/models"
	"github.com/AutoDruid/photon-parser"
)

func DecodeUpdateFame(params []photon.ParameterV18) (models.UpdateFameEvent, bool) {
	var e models.UpdateFameEvent

	log.Println("Decode Update Fame", params)

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
