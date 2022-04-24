package event

import (
	"reflect"
)

type Event interface{}

type Handler func(Event) error

type Publisher interface {
	Publish(Event) error
}

type Registrar interface {
	RegisterEventHandler(e Event, handler Handler)
}

type Bus struct {
	handlers map[reflect.Type][]Handler
}

func NewBus() *Bus {
	return &Bus{
		handlers: make(map[reflect.Type][]Handler),
	}
}

func (b *Bus) RegisterEventHandler(e Event, handler Handler) {
	t := reflect.TypeOf(e)

	handlers, ok := b.handlers[t]
	if !ok {
		handlers = make([]Handler, 0, 1)
	}

	b.handlers[t] = append(handlers, handler)
}

func (b *Bus) Publish(e Event) error {
	t := reflect.TypeOf(e)

	handlers, ok := b.handlers[t]
	if !ok {
		return nil
	}

	for _, handler := range handlers {
		if err := handler(e); err != nil {
			return err
		}
	}

	return nil
}
