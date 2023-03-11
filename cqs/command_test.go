package cqs_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/kit/cqs"
)

func TestCommandHandlerMultiMiddleware(t *testing.T) {
	require := require.New(t)

	var mwExecutionOrder []int

	commandHandlerMw := func(count int) cqs.CommandHandlerMiddleware[cqs.Command] {
		return func(h cqs.CommandHandler[cqs.Command]) cqs.CommandHandler[cqs.Command] {
			return &CommandHandlerMock[cqs.Command]{
				HandleFunc: func(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
					_, _ = h.Handle(ctx, cmd)
					mwExecutionOrder = append(mwExecutionOrder, count)

					return nil, nil
				},
			}
		}
	}

	multiMw := cqs.CommandHandlerMultiMiddleware(
		commandHandlerMw(1),
		commandHandlerMw(2),
		commandHandlerMw(3),
	)

	handlerExecutionCount := 0
	ch := &CommandHandlerMock[cqs.Command]{
		HandleFunc: func(context.Context, cqs.Command) ([]cqs.Event, error) {
			handlerExecutionCount++
			return nil, nil
		},
	}

	_, err := multiMw(ch).Handle(context.Background(), nil)
	require.NoError(err)
	require.Equal([]int{1, 2, 3}, mwExecutionOrder)
	require.Equal(1, handlerExecutionCount)
}
