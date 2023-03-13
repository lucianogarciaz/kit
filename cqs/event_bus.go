package cqs

import (
	"context"
	"errors"
	"fmt"
)

const errMsgBus = "dispatch event: %w"

var (
	ErrEmptyEventName          = errors.New("empty event name")
	ErrEmptyEventHandler       = errors.New("empty event handler")
	ErrEmptyEventToCommandFunc = errors.New("empty event to command func")
	ErrEmptyCommandHandler     = errors.New("empty command handler")
)

// EventsBusError is self-described.
type EventsBusError struct {
	error
}

// Unwrap returns the underlying error.
func (e EventsBusError) Unwrap() error {
	return errors.Unwrap(e.error)
}

func newEventsBusError(err error) error {
	return EventsBusError{fmt.Errorf(errMsgBus, err)}
}

// EventHandler is self-explanatory.
type EventHandler interface {
	Handle(ctx context.Context, event Event) error
}

// An EventsBus is the piece that relates events and its handlers.
type EventsBus interface {
	Subscribe(name EventName, handler EventHandler) error
	Dispatch(ctx context.Context, event Event) error
}

// BasicEventsBus is not concurrent-safe for subscribing and dispatching at the same time
// so its handlers should be subscribed first before dispatching events.
type BasicEventsBus struct {
	handlersByName map[EventName][]EventHandler
}

var _ EventsBus = &BasicEventsBus{}

// Subscribe links a specific event with its handler.
func (bus *BasicEventsBus) Subscribe(name EventName, handler EventHandler) error {
	if name == "" {
		return ErrEmptyEventName
	}

	if handler == nil {
		return ErrEmptyEventHandler
	}

	if bus.handlersByName == nil {
		bus.handlersByName = make(map[EventName][]EventHandler)
	}

	if _, ok := bus.handlersByName[name]; !ok {
		bus.handlersByName[name] = make([]EventHandler, 0)
	}

	bus.handlersByName[name] = append(bus.handlersByName[name], handler)

	return nil
}

// Dispatch receives an event and calls its handler. Retries must be handled by the caller.
func (bus BasicEventsBus) Dispatch(ctx context.Context, ev Event) error {
	hs, ok := bus.handlersByName[ev.EventName()]
	if !ok {
		return nil
	}

	multierror := NewMultiError()

	for _, h := range hs {
		if err := h.Handle(ctx, ev); err != nil {
			multierror.Add(newEventsBusError(err))
		}
	}

	return multierror.ErrResult()
}

// EventToCommandFunc is a function to convert each event to its correspondent command.
type EventToCommandFunc func(ev Event) (Command, error)

var _ EventHandler = &EventCommandHandler{}

// EventCommandHandler is a type for linking an event to its resultant commandHandler.
type EventCommandHandler struct {
	commandHandler     CommandHandler[Command]
	eventToCommandFunc EventToCommandFunc
}

// Handle calls the underlying commandHandler.
func (h EventCommandHandler) Handle(ctx context.Context, ev Event) error {
	cmd, err := h.eventToCommandFunc(ev)
	if err != nil {
		return err
	}

	_, err = h.commandHandler.Handle(ctx, cmd)

	return err
}

// NewEventCommandHandler is self-explanatory.
func NewEventCommandHandler(f EventToCommandFunc, ch CommandHandler[Command]) (EventCommandHandler, error) {
	if f == nil {
		return EventCommandHandler{}, ErrEmptyEventToCommandFunc
	}

	if ch == nil {
		return EventCommandHandler{}, ErrEmptyCommandHandler
	}

	return EventCommandHandler{
		commandHandler:     ch,
		eventToCommandFunc: f,
	}, nil
}
