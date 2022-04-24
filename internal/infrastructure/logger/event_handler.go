package logger

import (
	"fmt"
	"stabulum/internal/common/event"
	"stabulum/internal/common/logger"
	"stabulum/internal/domain/product"
)

type EventHandler struct {
	logger logger.Logger
}

func NewEventHandler(logger logger.Logger) *EventHandler {
	return &EventHandler{
		logger: logger,
	}
}

func (h *EventHandler) RegisterEvents(eventRegistrar event.Registrar) {
	eventRegistrar.RegisterEventHandler(product.CreatedEvent{}, h.handleProductCreatedEvent)
}

func (h *EventHandler) handleProductCreatedEvent(e event.Event) error {
	productCreatedEvent, ok := e.(product.CreatedEvent)
	if !ok {
		return nil
	}

	h.logger.Println(fmt.Sprintf("product %s created", productCreatedEvent.Name))

	return nil
}
