package cqs_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/kit/cqs"
)

func TestEventBusSubscribe(t *testing.T) {
	require := require.New(t)

	eventName := cqs.EventName("foo")

	t.Run(`Given an event bus,
	when Subscribe is called with an empty event name,
	then it returns an ErrEmptyEventName`, func(t *testing.T) {
		bus := cqs.BasicEventsBus{}

		err := bus.Subscribe("", nil)
		require.ErrorIs(err, cqs.ErrEmptyEventName)
	})

	t.Run(`Given an event bus,
	when Subscribe is called with an empty event handler,
	then it returns an ErrEmptyEventHandler`, func(t *testing.T) {
		bus := cqs.BasicEventsBus{}

		err := bus.Subscribe(eventName, nil)
		require.ErrorIs(err, cqs.ErrEmptyEventHandler)
	})

	t.Run(`Given an event bus and an event handler,
	when Subscribe is called,
	then it subscribed correctly`, func(t *testing.T) {
		bus := cqs.BasicEventsBus{}
		evHandlerMock1 := &EventHandlerMock{
			HandleFunc: nil,
		}

		err := bus.Subscribe(eventName, evHandlerMock1)
		require.NoError(err)
	})

	t.Run(`Given an event bus and an event handler different from the subscribed,
	when Subscribe is called with the same event name,
	then it subscribed correctly`, func(t *testing.T) {
		bus := cqs.BasicEventsBus{}
		evHandlerMock1 := &EventHandlerMock{
			HandleFunc: nil,
		}

		err := bus.Subscribe(eventName, evHandlerMock1)
		require.NoError(err)

		evHandlerMock2 := &EventHandlerMock{
			HandleFunc: nil,
		}

		err = bus.Subscribe(eventName, evHandlerMock2)
		require.NoError(err)
	})
}

func TestEventBusDispatch(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()
	eventName := cqs.EventName("foo")

	t.Run(`Given an event bus,
	when Dispatch is called without corresponding handler,
	then it does nothing`, func(t *testing.T) {
		bus := cqs.BasicEventsBus{}

		ev := &EventMock{
			EventNameFunc: func() cqs.EventName { return eventName },
		}

		err := bus.Dispatch(ctx, ev)
		require.NoError(err)
	})

	t.Run(`Given an event bus and an event handler that returns error on Handle,
	when Dispatch is called,
	then it returns error`, func(t *testing.T) {
		expectedError := errors.New("event handler error")
		bus := cqs.BasicEventsBus{}
		evHandlerMock1 := &EventHandlerMock{
			HandleFunc: func(context.Context, cqs.Event) error {
				return expectedError
			},
		}

		err := bus.Subscribe(eventName, evHandlerMock1)
		require.NoError(err)

		ev := &EventMock{
			EventNameFunc: func() cqs.EventName { return eventName },
		}

		err = bus.Dispatch(ctx, ev)
		require.ErrorContains(err, expectedError.Error())
	})

	t.Run(`Given an event bus and an event handler,
	when Dispatch is called
	then it returns no error and the event handlers are called`, func(t *testing.T) {
		bus := cqs.BasicEventsBus{}
		evHandlerMock1 := &EventHandlerMock{
			HandleFunc: func(context.Context, cqs.Event) error {
				return nil
			},
		}

		err := bus.Subscribe(eventName, evHandlerMock1)
		require.NoError(err)

		evHandlerMock2 := &EventHandlerMock{
			HandleFunc: func(context.Context, cqs.Event) error {
				return nil
			},
		}

		err = bus.Subscribe(eventName, evHandlerMock2)
		require.NoError(err)

		ev := &EventMock{
			EventNameFunc: func() cqs.EventName { return eventName },
		}

		err = bus.Dispatch(ctx, ev)
		require.NoError(err)

		eventHandlerMockCalls := evHandlerMock1.HandleCalls()
		require.Len(eventHandlerMockCalls, 1)
		require.Equal(ev, eventHandlerMockCalls[0].Event)
		eventHandlerMockCalls = evHandlerMock2.HandleCalls()
		require.Len(eventHandlerMockCalls, 1)
		require.Equal(ev, eventHandlerMockCalls[0].Event)
	})
}

func TestNewEventCommandHandler(t *testing.T) {
	require := require.New(t)

	t.Run(`Given a event command handler with an empty eventToCommandFunc,
	when it's called,
	then it returns error`, func(t *testing.T) {
		cmdHandlerMock := &CommandHandlerMock[cqs.Command]{}
		eventCommandHandler, err := cqs.NewEventCommandHandler(nil, cmdHandlerMock)
		require.Empty(eventCommandHandler)
		require.Equal(err, cqs.ErrEmptyEventToCommandFunc)
	})

	t.Run(`Given a event command handler with an empty command handler,
	when it's called,
	then it returns error`, func(t *testing.T) {
		eventToCommandFunc := func(e cqs.Event) (cqs.Command, error) { return nil, nil }
		eventCommandHandler, err := cqs.NewEventCommandHandler(eventToCommandFunc, nil)
		require.ErrorIs(err, cqs.ErrEmptyCommandHandler)
		require.Empty(eventCommandHandler)
	})

	t.Run(`Given a event command handler,
	when it's called,
	then it returns no error`, func(t *testing.T) {
		cmdHandlerMock := &CommandHandlerMock[cqs.Command]{}
		eventToCommandFunc := func(e cqs.Event) (cqs.Command, error) { return nil, nil }

		eventCH, err := cqs.NewEventCommandHandler(eventToCommandFunc, cmdHandlerMock)
		require.NoError(err)
		require.NotEmpty(eventCH)
	})
}

func TestEventCommandHandlerHandle(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()
	eventName := cqs.EventName("foo")

	t.Run(`Given an EventCommandHandler with an eventToCommandFunc that returns error,
	when it's called,
	then it returns the error`, func(t *testing.T) {
		expectedError := errors.New("eventToCommandError")
		eventToCommandFunc := func(e cqs.Event) (cqs.Command, error) { return nil, expectedError }
		cmdHandlerMock := &CommandHandlerMock[cqs.Command]{}
		eventCH, err := cqs.NewEventCommandHandler(eventToCommandFunc, cmdHandlerMock)
		require.NoError(err)
		evMock := &EventMock{
			EventNameFunc: func() cqs.EventName { return eventName },
		}
		require.ErrorIs(eventCH.Handle(ctx, evMock), expectedError)
	})

	t.Run(`Given an EventCommandHandler,
	when it's called,
	then it returns the error`, func(t *testing.T) {
		cmdMock := &CommandMock{}
		eventToCommandFunc := func(e cqs.Event) (cqs.Command, error) { return cmdMock, nil }
		cmdHandlerMock := &CommandHandlerMock[cqs.Command]{
			HandleFunc: func(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
				return nil, nil
			},
		}

		eventCH, err := cqs.NewEventCommandHandler(eventToCommandFunc, cmdHandlerMock)
		require.NoError(err)

		evMock := &EventMock{
			EventNameFunc: func() cqs.EventName { return eventName },
		}
		require.NoError(eventCH.Handle(ctx, evMock))
		require.Len(cmdHandlerMock.HandleCalls(), 1)
		require.Equal(cmdHandlerMock.HandleCalls()[0].Cmd, cmdMock)
	})
}
