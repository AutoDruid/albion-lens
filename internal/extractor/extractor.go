package extractor

import (
	"github.com/AutoDruid/photon-parser"
)

type Kind uint8

const (
	KindEvent     Kind = 252
	KindOperation Kind = 253
)

type RawOperation struct {
	Code int
	Kind
	Params []photon.ParameterV18
}

func Extract(params []photon.ParameterV18) (RawOperation, bool) {
	var op RawOperation
	var payload []photon.ParameterV18

	for _, p := range params {
		if p.ID() == uint8(KindEvent) || p.ID() == uint8(KindOperation) {
			v, ok := p.IntValue()
			if !ok {
				return op, false
			}
			op.Code = int(v)
			op.Kind = Kind(p.ID())
		} else {
			payload = append(payload, p)
		}
	}

	op.Params = payload
	return op, op.Code != 0
}
