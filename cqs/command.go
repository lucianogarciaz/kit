package cqs

import (
	"context"
)

// Command is the interface for identifying commands by name.
type Command interface {
	CommandName() string
}

// CommandHandler is the interface for implementing the logic that mutates our domain.
type CommandHandler[C Command] interface {
	Handle(ctx context.Context, cmd C) ([]Event, error)
}

// CommandHandlerFunc is a function that implements CommandHandler interface.
type CommandHandlerFunc[C Command] func(ctx context.Context, cmd C) ([]Event, error)

// Handle is the CommandHandler interface implementation.
func (f CommandHandlerFunc[C]) Handle(ctx context.Context, cmd C) ([]Event, error) {
	return f(ctx, cmd)
}

// CommandHandlerMiddleware is a type for decorating CommandHandlers.
type CommandHandlerMiddleware[C Command] func(h CommandHandler[C]) CommandHandler[C]

// CommandHandlerMultiMiddleware applies a sequence of middlewares to a given command handler.
func CommandHandlerMultiMiddleware[C Command](middlewares ...CommandHandlerMiddleware[C]) CommandHandlerMiddleware[C] {
	return func(h CommandHandler[C]) CommandHandler[C] {
		handler := h
		for _, m := range middlewares {
			handler = m(handler)
		}

		return CommandHandlerFunc[C](handler.Handle)
	}
}
