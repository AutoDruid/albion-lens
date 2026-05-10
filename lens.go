package albionlens

import (
	"github.com/AutoDruid/albion-lens/internal/decoders"
	"github.com/AutoDruid/albion-lens/internal/extractor"
	"github.com/AutoDruid/albion-lens/models"
	eventTypes "github.com/AutoDruid/albion-lens/types/events"
	operationTypes "github.com/AutoDruid/albion-lens/types/operations"
	"github.com/AutoDruid/photon-parser"
)

type Lens struct {
	parser     *photon.Parser[photon.ParameterV18]
	events     map[eventTypes.EventCode]func([]photon.ParameterV18)
	operations map[operationTypes.OperationCode]func([]photon.ParameterV18)
}

func NewLens() *Lens {
	l := &Lens{
		parser:     photon.NewV18(),
		events:     make(map[eventTypes.EventCode]func([]photon.ParameterV18)),
		operations: make(map[operationTypes.OperationCode]func([]photon.ParameterV18)),
	}
	l.startListeners()
	return l
}

func (l *Lens) Parse(data []byte) error {
	if _, err := l.parser.ParsePacket(data); err != nil {
		return err
	}
	return nil
}

func (l *Lens) startListeners() {
	l.parser.OnEvent(func(reliable photon.ReliableV18) {
		op, ok := extractor.Extract(reliable.Parameters)
		if !ok {
			return
		}

		switch op.Kind {
		case extractor.KindEvent:
			if fn, ok := l.events[eventTypes.EventCode(op.Code)]; ok {
				fn(op.Params)
			}
		case extractor.KindOperation:
			if fn, ok := l.operations[operationTypes.OperationCode(op.Code)]; ok {
				fn(op.Params)
			}
		}
	})
}

func (l *Lens) OnMove(fn func(models.MoveEvent)) {
	l.operations[operationTypes.Move] = func(params []photon.ParameterV18) {
		e, ok := decoders.DecodeMove(params)
		if !ok {
			return
		}
		fn(e)
	}
}

func (l *Lens) OnUpdateFame(fn func(models.UpdateFameEvent)) {
	l.events[eventTypes.UpdateFame] = func(params []photon.ParameterV18) {
		e, ok := decoders.DecodeUpdateFame(params)
		if !ok {
			return
		}
		fn(e)
	}
}

func (l *Lens) OnCustomEvents(eventCode eventTypes.EventCode, fn func([]photon.ParameterV18)) {
	l.events[eventCode] = func(params []photon.ParameterV18) {
		fn(params)
	}
}

func (l *Lens) OnCustomOperations(operationCode operationTypes.OperationCode, fn func([]photon.ParameterV18)) {
	l.operations[operationCode] = func(params []photon.ParameterV18) {
		fn(params)
	}
}
