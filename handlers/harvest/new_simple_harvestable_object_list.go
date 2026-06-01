package harvest

import (
	"log"

	albionlens "github.com/AutoDruid/albion-lens"
	"github.com/AutoDruid/albion-lens/handlers"
	"github.com/AutoDruid/albion-lens/types/events"
	"github.com/AutoDruid/photon-parser"
)

type NewSimpleHarvestableObjectListEvent struct {
	ClusterObjectIDs []uint16
	Type             []uint8
	Tiers            []uint8
	Positions        []handlers.Vec2
	ChargeAvailable  []uint8
}

func decodeNewSimpleHarvestableObjectList(params []photon.ParameterV18) (NewSimpleHarvestableObjectListEvent, bool) {
	var e NewSimpleHarvestableObjectListEvent

	size := int(params[0].Value.Num)

	e.ClusterObjectIDs = make([]uint16, size)
	for i, value := range params[0].Int16ArrayValue() {
		log.Println("value", i, value)
		e.ClusterObjectIDs[i] = uint16(value)
	}

	e.Type = make([]uint8, size)
	for i, value := range params[1].ByteArrayValue() {
		e.Type[i] = uint8(value)
	}

	e.Tiers = make([]uint8, size)
	for i, value := range params[2].ByteArrayValue() {
		e.Tiers[i] = uint8(value)
	}

	e.Positions = make([]handlers.Vec2, size)
	temp := handlers.Collect(params[3].Float32ArrayValue(), uint64(size*2))
	for i := 0; i < size; i += 2 {
		e.Positions[i].X = temp[i]
		e.Positions[i].Y = temp[i+1]
	}

	e.ChargeAvailable = make([]uint8, size)
	for i, value := range params[4].ByteArrayValue() {
		e.ChargeAvailable[i] = uint8(value)
	}

	return e, true
}

var NewSimpleHarvestableObjectListHandler = albionlens.EventHandler[NewSimpleHarvestableObjectListEvent]{
	Code:    events.NewSimpleHarvestableObjectList,
	Handler: decodeNewSimpleHarvestableObjectList,
}
