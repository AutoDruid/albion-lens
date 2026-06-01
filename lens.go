package albionlens

import (
	"github.com/AutoDruid/albion-lens/internal"
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
		parser:     photon.NewParserV18(),
		events:     make(map[eventTypes.EventCode]func([]photon.ParameterV18)),
		operations: make(map[operationTypes.OperationCode]func([]photon.ParameterV18)),
	}
	l.startListeners()
	return l
}

func (l *Lens) Parse(data []byte) error {
	var session photon.SessionV18
	if err := l.parser.ParsePacketInto(data, &session); err != nil {
		return err
	}
	return nil
}

func (l *Lens) startListeners() {

	l.parser.OnEventData(func(reliable photon.ReliableV18) {
		op, ok := internal.Extract(reliable.Parameters)
		if !ok {
			return
		}
		if fn, ok := l.events[eventTypes.EventCode(op.Code)]; ok {
			fn(reliable.Parameters)
		}
	})

	l.parser.OnOperationResponse(func(reliable photon.ReliableV18) {
		op, ok := internal.Extract(reliable.Parameters)
		if !ok {
			return
		}
		if fn, ok := l.operations[operationTypes.OperationCode(op.Code)]; ok {
			fn(reliable.Parameters)
		}
	})
}

type Code interface {
	eventTypes.EventCode | operationTypes.OperationCode
}

type Handler[E any, C Code] struct {
	Code    C
	Handler func([]photon.ParameterV18) (E, bool)
}

type EventHandler[E any] = Handler[E, eventTypes.EventCode]

func OnEvent[E any](l *Lens, h EventHandler[E], cb func(E)) {
	l.events[h.Code] = func(p []photon.ParameterV18) {
		if e, ok := h.Handler(p); ok {
			cb(e)
		}
	}
}

func OnCustomEvent(l *Lens, key eventTypes.EventCode, cb func([]photon.ParameterV18)) {
	l.events[key] = func(p []photon.ParameterV18) {
		cb(p)
	}
}

type OperationHandler[E any] = Handler[E, operationTypes.OperationCode]

func OnOperation[E any](l *Lens, h OperationHandler[E], cb func(E)) {
	l.operations[h.Code] = func(p []photon.ParameterV18) {
		if e, ok := h.Handler(p); ok {
			cb(e)
		}
	}
}

func OnCustomOperation(l *Lens, key operationTypes.OperationCode, cb func([]photon.ParameterV18)) {
	l.operations[key] = func(p []photon.ParameterV18) {
		cb(p)
	}
}
